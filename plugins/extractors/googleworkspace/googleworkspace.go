package googleworkspace

import (
	"context"
	_ "embed" // used to print the embedded assets
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"

	"github.com/odpf/meteor/models"
	"github.com/odpf/meteor/plugins"
	"github.com/odpf/meteor/registry"
	"github.com/odpf/meteor/utils"
	"github.com/odpf/salt/log"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"

	commonv1beta1 "github.com/odpf/meteor/models/odpf/assets/common/v1beta1"
	facetsv1beta1 "github.com/odpf/meteor/models/odpf/assets/facets/v1beta1"
	assetsv1beta1 "github.com/odpf/meteor/models/odpf/assets/v1beta1"
)

//go:embed README.md
var summary string

type Config struct {
	ServiceAccountJSON string `mapstructure:"service_account_json" validate:"required"`
	UserEmail          string `mapstructure:"user_email" validate:"required"`
}

var sampleConfig = `
service_account_json: {
    "type": "service_account",
    "project_id": "odpf-project",
    "private_key_id": "3cb2saasa3ef788dvdvdvdvdvdssdvds57",
    "private_key": "-----BEGIN PRIVATE KEY-----\njbjabdjbajd\n-----END PRIVATE KEY-----\n",
    "client_email": "meteor-sa@odpf-project.iam.gserviceaccount.com",
    "client_id": "1100599572858548635286",
    "auth_uri": "https://accounts.google.com/o/oauth2/auth",
    "token_uri": "https://oauth2.googleapis.com/token",
    "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
    "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/meteor-sa%40odpf-project.iam.gserviceaccount.com"
}
user_email: user@odpf.com`

var info = plugins.Info{
	Description:  "User list from Google Workspace",
	SampleConfig: sampleConfig,
	Tags:         []string{"platform", "extractor"},
	Summary:      summary,
}

// Extractor manages the extraction of data from the extractor
type Extractor struct {
	plugins.BaseExtractor
	logger      log.Logger
	config      Config
	TokenSource oauth2.TokenSource
	Client      *admin.Service
	emit        plugins.Emit
}

// New returns a pointer to an initialized Extractor Object
func New(logger log.Logger) *Extractor {
	e := &Extractor{
		logger: logger,
	}
	e.BaseExtractor = plugins.NewBaseExtractor(info, &e.config)

	return e
}

// Init initializes the extractor
func (e *Extractor) Init(ctx context.Context, config plugins.Config) (err error) {
	if err = e.BaseExtractor.Init(ctx, config); err != nil {
		return err
	}

	jwtConfig, err := google.JWTConfigFromJSON([]byte(e.config.ServiceAccountJSON), admin.AdminDirectoryUserScope)
	if err != nil {
		return plugins.InvalidConfigError{}
	}

	if e.config.UserEmail == "" {
		return plugins.InvalidConfigError{}
	}
	jwtConfig.Subject = e.config.UserEmail
	ts := jwtConfig.TokenSource(ctx)

	e.TokenSource = ts

	return
}

// Extract extracts the data from the extractor
// The data is returned as a list of assets.Asset
func (e *Extractor) Extract(ctx context.Context, emit plugins.Emit) (err error) {
	var status string
	e.emit = emit
	r, err := FetchUsers(ctx, e.TokenSource)
	if err != nil {
		return err
	}

	if len(r.Users) == 0 {
		e.logger.Info("No users found.\n")
		return nil
	} else {
		for _, u := range r.Users {
			if !u.Suspended {
				status = "not suspended"
			} else {
				status = "suspended"
			}

			var userAttributes = make(map[string]interface{})
			userAttributes = buildOrgAttributes(userAttributes, u.Organizations)
			userAttributes = buildRelationsAttributes(userAttributes, u.Relations)
			userAttributes = buildCustomSchemasAttributes(userAttributes, u.CustomSchemas)
			userAttributes["org_unit_path"] = u.OrgUnitPath
			userAttributes["aliases"] = strings.Join(u.Aliases, ",")

			e.emit(models.NewRecord(&assetsv1beta1.User{
				Resource: &commonv1beta1.Resource{
					Service: "google workspace",
					Name:    u.Name.FullName,
					Urn:     u.PrimaryEmail,
				},
				Email:    u.PrimaryEmail,
				FullName: u.Name.FullName,
				Status:   status,
				Properties: &facetsv1beta1.Properties{
					Attributes: utils.TryParseMapToProto(userAttributes),
				},
			}))
		}
	}

	return nil
}

// init registers the extractor to catalog
func init() {
	if err := registry.Extractors.Register("googleworkspace", func() plugins.Extractor {
		return New(plugins.GetLog())
	}); err != nil {
		panic(err)
	}
}

func FetchUsers(ctx context.Context, ts oauth2.TokenSource) (*admin.Users, error) {
	srv, err := admin.NewService(ctx, option.WithTokenSource(ts))
	if err != nil {
		return nil, err
	}

	r, err := srv.Users.List().Customer("my_customer").Do()
	if err != nil {
		return nil, errors.Wrap(err, "Unable to retrieve users in domain")
	}

	return r, nil
}

func buildOrgAttributes(userAttributes map[string]interface{}, orgsIfc interface{}) map[string]interface{} {
	if orgsIfc != nil {
		orgs := reflect.ValueOf(orgsIfc)
		if orgs.Kind() == reflect.Slice {
			//based on assumpton that a user shall belong to a single org
			org := reflect.ValueOf(orgs.Index(0).Interface())
			for _, key := range org.MapKeys() {
				value := org.MapIndex(key)
				userAttributes[fmt.Sprintf("%v", key.Interface())] = value.Interface()
			}
		}
	}
	return userAttributes
}

func buildRelationsAttributes(userAttributes map[string]interface{}, relationsIfc interface{}) map[string]interface{} {
	if relationsIfc != nil {
		relations := reflect.ValueOf(relationsIfc)
		if relations.Kind() == reflect.Slice {
			for idx := 0; idx < relations.Len(); idx++ {
				var relationType, relationValue string

				relation := reflect.ValueOf(relations.Index(idx).Interface())
				for _, key := range relation.MapKeys() {
					value := relation.MapIndex(key)

					if key.Interface().(string) == "type" {
						relationType = value.Interface().(string)
					} else if key.Interface().(string) == "value" {
						relationValue = value.Interface().(string)
					}
				}
				userAttributes[relationType] = relationValue
			}
		}
	}
	return userAttributes
}

func buildCustomSchemasAttributes(userAttributes map[string]interface{}, customSchemasIfc interface{}) map[string]interface{} {
	if customSchemasIfc != nil {
		customSchema := reflect.ValueOf(customSchemasIfc)
		if customSchema.Kind() == reflect.Map {
			for _, key := range customSchema.MapKeys() {
				value := customSchema.MapIndex(key)
				userAttributes[fmt.Sprintf("%v", key.Interface())] = value.Interface()
			}
		}
	}
	return userAttributes
}

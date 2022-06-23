"use strict";(self.webpackChunkmeteor=self.webpackChunkmeteor||[]).push([[840],{3905:(e,t,r)=>{r.d(t,{Zo:()=>p,kt:()=>m});var n=r(7294);function a(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function i(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function o(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?i(Object(r),!0).forEach((function(t){a(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function s(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},i=Object.keys(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var l=n.createContext({}),c=function(e){var t=n.useContext(l),r=t;return e&&(r="function"==typeof e?e(t):o(o({},t),e)),r},p=function(e){var t=c(e.components);return n.createElement(l.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},d=n.forwardRef((function(e,t){var r=e.components,a=e.mdxType,i=e.originalType,l=e.parentName,p=s(e,["components","mdxType","originalType","parentName"]),d=c(r),m=a,g=d["".concat(l,".").concat(m)]||d[m]||u[m]||i;return r?n.createElement(g,o(o({ref:t},p),{},{components:r})):n.createElement(g,o({ref:t},p))}));function m(e,t){var r=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=r.length,o=new Array(i);o[0]=d;var s={};for(var l in t)hasOwnProperty.call(t,l)&&(s[l]=t[l]);s.originalType=e,s.mdxType="string"==typeof e?e:a,o[1]=s;for(var c=2;c<i;c++)o[c]=r[c];return n.createElement.apply(null,o)}return n.createElement.apply(null,r)}d.displayName="MDXCreateElement"},4953:(e,t,r)=>{r.r(t),r.d(t,{contentTitle:()=>o,default:()=>p,frontMatter:()=>i,metadata:()=>s,toc:()=>l});var n=r(7462),a=(r(7294),r(3905));const i={},o="Plugins",s={unversionedId:"guides/list_Plugins",id:"guides/list_Plugins",isDocsHomePage:!1,title:"Plugins",description:"Before getting started we expect you went through the prerequisites.",source:"@site/docs/guides/1_list_Plugins.md",sourceDirName:"guides",slug:"/guides/list_Plugins",permalink:"/meteor/docs/guides/list_Plugins",editUrl:"https://github.com/odpf/meteor/edit/master/docs/docs/guides/1_list_Plugins.md",tags:[],version:"current",sidebarPosition:1,frontMatter:{},sidebar:"docsSidebar",previous:{title:"Installation",permalink:"/meteor/docs/guides/installation"},next:{title:"Recipes - Creation and linting",permalink:"/meteor/docs/guides/manage_recipes"}},l=[{value:"Listing all the plugins",id:"listing-all-the-plugins",children:[]}],c={toc:l};function p(e){let{components:t,...r}=e;return(0,a.kt)("wrapper",(0,n.Z)({},c,r,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"plugins"},"Plugins"),(0,a.kt)("p",null,"Before getting started we expect you went through the ",(0,a.kt)("a",{parentName:"p",href:"/meteor/docs/guides/introduction#prerequisites"},"prerequisites"),"."),(0,a.kt)("p",null,"Meteor follows a plugin driven approach and hence includes basically three kinds of plugins for the metadata orchestration: extractors (source), processors, and sinks (destination).\nSome details on these 3 are:"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("p",{parentName:"li"},(0,a.kt)("strong",{parentName:"p"},"Extractors")," are the set of plugins that are the source of our metadata and include databases, dashboards, users, etc.")),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("p",{parentName:"li"},(0,a.kt)("strong",{parentName:"p"},"Processors")," are the set of plugins that perform the enrichment or data processing for the metadata after extraction.")),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("p",{parentName:"li"},(0,a.kt)("strong",{parentName:"p"},"Sinks")," are the plugins that act as the destination of our metadata after extraction and processing."))),(0,a.kt)("p",null,"Read more about the concepts on each of these in ",(0,a.kt)("a",{parentName:"p",href:"/meteor/docs/concepts/overview"},"concepts"),".\nTo get more context on these plugins, it is recommended to try out the ",(0,a.kt)("inlineCode",{parentName:"p"},"list")," command to get the list of plugins of a specific type. Commands to list the plugins are mentioned below"),(0,a.kt)("h2",{id:"listing-all-the-plugins"},"Listing all the plugins"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"# list all available extractors\n$ meteor list extractors\n\n# list all extractors with alias 'e'\n$ meteor list e\n\n# list available sinks\n$ meteor list sinks\n\n# list all sinks with alias 's'\n$ meteor list s\n\n# list all available processors\n$ meteor list processors\n\n# list all processors with alias 'p'\n$ meteor list p\n")))}p.isMDXComponent=!0}}]);
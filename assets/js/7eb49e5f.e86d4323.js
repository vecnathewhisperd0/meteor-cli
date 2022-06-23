"use strict";(self.webpackChunkmeteor=self.webpackChunkmeteor||[]).push([[242],{3905:(e,t,r)=>{r.d(t,{Zo:()=>p,kt:()=>m});var n=r(7294);function o(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function a(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function s(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?a(Object(r),!0).forEach((function(t){o(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):a(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function c(e,t){if(null==e)return{};var r,n,o=function(e,t){if(null==e)return{};var r,n,o={},a=Object.keys(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||(o[r]=e[r]);return o}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(o[r]=e[r])}return o}var i=n.createContext({}),l=function(e){var t=n.useContext(i),r=t;return e&&(r="function"==typeof e?e(t):s(s({},t),e)),r},p=function(e){var t=l(e.components);return n.createElement(i.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},d=n.forwardRef((function(e,t){var r=e.components,o=e.mdxType,a=e.originalType,i=e.parentName,p=c(e,["components","mdxType","originalType","parentName"]),d=l(r),m=o,f=d["".concat(i,".").concat(m)]||d[m]||u[m]||a;return r?n.createElement(f,s(s({ref:t},p),{},{components:r})):n.createElement(f,s({ref:t},p))}));function m(e,t){var r=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var a=r.length,s=new Array(a);s[0]=d;var c={};for(var i in t)hasOwnProperty.call(t,i)&&(c[i]=t[i]);c.originalType=e,c.mdxType="string"==typeof e?e:o,s[1]=c;for(var l=2;l<a;l++)s[l]=r[l];return n.createElement.apply(null,s)}return n.createElement.apply(null,r)}d.displayName="MDXCreateElement"},2608:(e,t,r)=>{r.r(t),r.d(t,{contentTitle:()=>s,default:()=>p,frontMatter:()=>a,metadata:()=>c,toc:()=>i});var n=r(7462),o=(r(7294),r(3905));const a={},s="Processor",c={unversionedId:"concepts/processor",id:"concepts/processor",isDocsHomePage:!1,title:"Processor",description:"A recipe can have none or many processors registered, depending upon the way the user wants metadata to be processed. A processor is basically a function that:",source:"@site/docs/concepts/processor.md",sourceDirName:"concepts",slug:"/concepts/processor",permalink:"/meteor/docs/concepts/processor",editUrl:"https://github.com/odpf/meteor/edit/master/docs/docs/concepts/processor.md",tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Source",permalink:"/meteor/docs/concepts/source"},next:{title:"Sink",permalink:"/meteor/docs/concepts/sink"}},i=[{value:"Built-in Processors",id:"built-in-processors",children:[{value:"metadata",id:"metadata",children:[]}]}],l={toc:i};function p(e){let{components:t,...r}=e;return(0,o.kt)("wrapper",(0,n.Z)({},l,r,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"processor"},"Processor"),(0,o.kt)("p",null,"A recipe can have none or many processors registered, depending upon the way the user wants metadata to be processed. A processor is basically a function that:"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},"expects a list of data"),(0,o.kt)("li",{parentName:"ul"},"processes the list"),(0,o.kt)("li",{parentName:"ul"},"returns a list")),(0,o.kt)("p",null,"The result from a processor will be passed on to the next processor until there is no more processor, hence the flow is sequential."),(0,o.kt)("h2",{id:"built-in-processors"},"Built-in Processors"),(0,o.kt)("h3",{id:"metadata"},"metadata"),(0,o.kt)("p",null,"This processor will set and overwrite metadata with given fields in the config."),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-yaml"},"processors:\n  - name: metadata\n    config:\n      fieldA: valueA\n      fieldB: valueB\n")),(0,o.kt)("table",null,(0,o.kt)("thead",{parentName:"table"},(0,o.kt)("tr",{parentName:"thead"},(0,o.kt)("th",{parentName:"tr",align:"left"},"key"),(0,o.kt)("th",{parentName:"tr",align:"left"},"Description"),(0,o.kt)("th",{parentName:"tr",align:"left"},"requirement"))),(0,o.kt)("tbody",{parentName:"table"},(0,o.kt)("tr",{parentName:"tbody"},(0,o.kt)("td",{parentName:"tr",align:"left"},(0,o.kt)("inlineCode",{parentName:"td"},"name")),(0,o.kt)("td",{parentName:"tr",align:"left"},"contains the name of processor"),(0,o.kt)("td",{parentName:"tr",align:"left"},"required")),(0,o.kt)("tr",{parentName:"tbody"},(0,o.kt)("td",{parentName:"tr",align:"left"},(0,o.kt)("inlineCode",{parentName:"td"},"config")),(0,o.kt)("td",{parentName:"tr",align:"left"},"different processors will require different config"),(0,o.kt)("td",{parentName:"tr",align:"left"},"required")))),(0,o.kt)("p",null,"More info about available processors can be found ",(0,o.kt)("a",{parentName:"p",href:"/meteor/docs/reference/processors"},"here"),"."))}p.isMDXComponent=!0}}]);
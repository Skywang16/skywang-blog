import{_ as b}from"./index.3a12671c.js";import{J as E,w as V}from"./element-plus.8c1ec78f.js";import P from"./menu.4639c37f.js";import k from"./preview.796add9c.js";import{_ as C}from"./attr-setting.vue_vue_type_script_setup_true_lang.ba9dbbe0.js";import{w as F}from"./index.d7de92a0.js";import{s as N,a as h}from"./decoration.cc6e9569.js";import{n as S,f as I,d as T}from"./index.a935f15b.js";import{d as D,a1 as A,r as _,e as f,w as J,an as O,o as g,c as U,X as i,P as u,a as M,u as p,k as v,Q as W,O as Q,U as R}from"./@vue.a137a740.js";import"./@vueuse.a2c70c0f.js";import"./@element-plus.3660753f.js";import"./lodash-es.7718a1e4.js";import"./dayjs.7a8df8b4.js";import"./axios.00a285e9.js";import"./async-validator.fb49d0f5.js";import"./@ctrl.fd318bfa.js";import"./escape-html.e5dfadb9.js";import"./normalize-wheel-es.8aeb3683.js";import"./lodash.1914dddf.js";import"./vue-router.9605b890.js";import"./pinia.706f3ea9.js";import"./css-color-function.1f5b2f4d.js";import"./color.52be626b.js";import"./clone.4f21711a.js";import"./color-convert.755d189f.js";import"./color-name.e7a4e1d3.js";import"./color-string.e356f5de.js";import"./balanced-match.d2a36341.js";import"./ms.564e106c.js";import"./nprogress.bcfbe0a4.js";import"./vue-clipboard3.76938e62.js";import"./clipboard.d4b72188.js";import"./echarts.098cf7d6.js";import"./zrender.0646a922.js";import"./tslib.60310f1a.js";import"./highlight.js.7165574c.js";import"./@highlightjs.7fc78ec7.js";import"./attr.vue_vue_type_script_setup_true_lang.fb3d0014.js";import"./index.ad18f60b.js";import"./picker.808d039d.js";import"./index.345a4530.js";import"./picker.781c9a4f.js";import"./index.33d8b440.js";import"./usePaging.f7aa3aa8.js";import"./index.vue_vue_type_script_setup_true_lang.7bfa50f8.js";import"./vue3-video-play.b1eef99b.js";import"./vuedraggable.6b9b01a3.js";import"./vue.b1bda70a.js";import"./sortablejs.57b1acee.js";import"./content.vue_vue_type_script_setup_true_lang.d5bbfe2c.js";import"./decoration-img.f6754dd5.js";import"./attr.vue_vue_type_script_setup_true_lang.2990574c.js";import"./content.29d83523.js";import"./attr.vue_vue_type_script_setup_true_lang.81637dbf.js";import"./add-nav.vue_vue_type_script_setup_true_lang.813c7320.js";import"./content.677ee7bf.js";import"./attr.vue_vue_type_script_setup_true_lang.231f3e9a.js";import"./content.vue_vue_type_script_setup_true_lang.719ecaea.js";import"./attr.vue_vue_type_script_setup_true_lang.a6ae21ee.js";import"./content.f83aca1e.js";import"./attr.vue_vue_type_script_setup_true_lang.9820b9a5.js";import"./content.b3504b05.js";import"./attr.vue_vue_type_script_setup_true_lang.04b2a4a7.js";import"./content.vue_vue_type_script_setup_true_lang.c0827591.js";import"./attr.vue_vue_type_script_setup_true_lang.9653ec08.js";import"./content.395cd987.js";const X={class:"decoration-pages min-w-[1100px]"},$={class:"flex h-full items-start"},j=D({name:"decorationPages"}),q=D({...j,setup(z){const m=e=>e.map(t=>{var s;return{id:S(),...((s=F[t])==null?void 0:s.options())||{}}}),o=A({[1]:{id:1,pageType:1,name:"\u9996\u9875\u88C5\u4FEE",pageData:m(["search","banner","nav","news"])},[2]:{id:2,pageType:2,name:"\u4E2A\u4EBA\u4E2D\u5FC3",pageData:m(["user-info","my-service","user-banner"])},[3]:{id:3,pageType:3,name:"\u5BA2\u670D\u8BBE\u7F6E",pageData:m(["customer-service"])}}),a=_("1"),r=_(-1),l=f(()=>{var e,t;return(t=(e=o[a.value])==null?void 0:e.pageData)!=null?t:[]}),x=f(()=>{var e,t;return(t=(e=o[a.value])==null?void 0:e.pageData[r.value])!=null?t:""}),c=async()=>{const e=await h({id:a.value});o[String(e.id)].pageData=JSON.parse(e.pageData)},w=async()=>{await N({...o[a.value],pageData:JSON.stringify(o[a.value].pageData)}),c(),I.msgSuccess("\u4FDD\u5B58\u6210\u529F")};return J(a,()=>{r.value=l.value.findIndex(e=>!e.disabled),c()},{immediate:!0}),(e,t)=>{const d=E,s=V,y=b,B=O("perms");return g(),U("div",X,[i(d,{shadow:"never",class:"!border-none flex-1 flex","body-style":{flex:1}},{default:u(()=>[M("div",$,[i(P,{modelValue:p(a),"onUpdate:modelValue":t[0]||(t[0]=n=>v(a)?a.value=n:null),menus:p(o)},null,8,["modelValue","menus"]),i(k,{modelValue:p(r),"onUpdate:modelValue":t[1]||(t[1]=n=>v(r)?r.value=n:null),pageData:p(l)},null,8,["modelValue","pageData"]),i(C,{class:"flex-1",widget:p(x)},null,8,["widget"])])]),_:1}),W((g(),Q(y,{class:"mt-4",fixed:!1},{default:u(()=>[i(s,{type:"primary",onClick:w},{default:u(()=>[R("\u4FDD\u5B58")]),_:1})]),_:1})),[[B,["decorate:pages:save"]]])])}}});const pt=T(q,[["__scopeId","data-v-a5fbac16"]]);export{pt as default};
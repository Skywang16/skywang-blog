import{C as q,D as O,Q,R as z,w as H,F as J,J as W,H as X,I as G,K as M}from"./element-plus.8c1ec78f.js";import{u as Y,_ as Z}from"./usePaging.f7aa3aa8.js";import{j as k,f as V,b as ee,_ as te}from"./index.a935f15b.js";import{d as P,a1 as D,b0 as oe,ag as ae,an as le,o as i,c as x,X as e,P as o,u as a,a9 as ne,W as se,a8 as ie,O as m,U as r,a as B,Q as g,V as L,T as re,k as ue}from"./@vue.a137a740.js";import{a as me,d as de,e as pe}from"./information.8c78de9c.js";import"./@vueuse.a2c70c0f.js";import"./@element-plus.3660753f.js";import"./lodash-es.7718a1e4.js";import"./dayjs.7a8df8b4.js";import"./axios.00a285e9.js";import"./async-validator.fb49d0f5.js";import"./@ctrl.fd318bfa.js";import"./escape-html.e5dfadb9.js";import"./normalize-wheel-es.8aeb3683.js";import"./lodash.1914dddf.js";import"./vue-router.9605b890.js";import"./pinia.706f3ea9.js";import"./css-color-function.1f5b2f4d.js";import"./color.52be626b.js";import"./clone.4f21711a.js";import"./color-convert.755d189f.js";import"./color-name.e7a4e1d3.js";import"./color-string.e356f5de.js";import"./balanced-match.d2a36341.js";import"./ms.564e106c.js";import"./nprogress.bcfbe0a4.js";import"./vue-clipboard3.76938e62.js";import"./clipboard.d4b72188.js";import"./echarts.098cf7d6.js";import"./zrender.0646a922.js";import"./tslib.60310f1a.js";import"./highlight.js.7165574c.js";import"./@highlightjs.7fc78ec7.js";const ce={class:"hero-lists"},_e={class:"flex justify-end mt-4"},fe=P({name:"heroLists"}),Ge=P({...fe,setup(he){const n=D({name:"",level:"",types:-1,status:-1,type:""}),{pager:d,getLists:c,resetPage:F,resetParams:A}=Y({fetchFun:pe,params:n}),_=D([]),v=async()=>{try{const l=await me();console.log(l,"response"),l.lists&&l.lists.length>0&&_.splice(0,_.length,...l.lists)}catch(l){console.error(l)}},T=async l=>{await V.confirm("\u786E\u5B9A\u8981\u5220\u9664\uFF1F"),await de({id:l}),V.msgSuccess("\u5220\u9664\u6210\u529F"),c()};return oe(()=>{v(),c()}),c(),v(),(l,s)=>{const U=q,f=O,h=Q,C=z,p=H,I=J,w=W,K=ee,E=ae("router-link"),u=X,N=te,R=G,$=Z,b=le("perms"),S=M;return i(),x("div",ce,[e(w,{class:"!border-none",shadow:"never"},{default:o(()=>[e(I,{ref:"formRef",class:"mb-[-16px]",model:a(n),inline:!0},{default:o(()=>[e(f,{label:"\u82F1\u96C4\u540D\u79F0"},{default:o(()=>[e(U,{class:"w-[280px]",modelValue:a(n).name,"onUpdate:modelValue":s[0]||(s[0]=t=>a(n).name=t),clearable:"",onKeyup:ne(a(F),["enter"])},null,8,["modelValue","onKeyup"])]),_:1}),e(f,{label:"\u82F1\u96C4\u72B6\u6001"},{default:o(()=>[e(C,{class:"w-[280px]",modelValue:a(n).status,"onUpdate:modelValue":s[1]||(s[1]=t=>a(n).status=t)},{default:o(()=>[e(h,{label:"\u5168\u90E8",value:-1}),e(h,{label:"\u663E\u793A",value:1}),e(h,{label:"\u9690\u85CF",value:0})]),_:1},8,["modelValue"])]),_:1}),e(f,{label:"\u82F1\u96C4\u5B9A\u4F4D"},{default:o(()=>[e(C,{class:"w-[280px]",modelValue:a(n).type,"onUpdate:modelValue":s[2]||(s[2]=t=>a(n).type=t)},{default:o(()=>[(i(!0),x(se,null,ie(a(_),t=>(i(),m(h,{key:t.id,label:t.name,value:t.id},null,8,["label","value"]))),128))]),_:1},8,["modelValue"])]),_:1}),e(f,null,{default:o(()=>[e(p,{type:"primary",onClick:a(F)},{default:o(()=>[r("\u67E5\u8BE2")]),_:1},8,["onClick"]),e(p,{onClick:a(A)},{default:o(()=>[r("\u91CD\u7F6E")]),_:1},8,["onClick"])]),_:1})]),_:1},8,["model"])]),_:1}),e(w,{class:"!border-none mt-4",shadow:"never"},{default:o(()=>[B("div",null,[g((i(),m(E,{to:{path:a(k)("hero:add/edit")}},{default:o(()=>[e(p,{type:"primary",class:"mb-4"},{icon:o(()=>[e(K,{name:"el-icon-Plus"})]),default:o(()=>[r(" \u65B0\u589E\u82F1\u96C4 ")]),_:1})]),_:1},8,["to"])),[[b,["hero:add","hero:add/edit"]]])]),g((i(),m(R,{size:"large",data:a(d).lists},{default:o(()=>[e(u,{label:"\u5E8F\u53F7",width:"60"},{default:o(t=>[r(L(t.$index+1),1)]),_:1}),e(u,{label:"ID",prop:"id","min-width":"80"}),e(u,{label:"\u56FE\u7247","min-width":"100"},{default:o(({row:t})=>[t.image?(i(),m(N,{key:0,src:t.image,width:60,height:45,"preview-src-list":[t.image],"preview-teleported":"",fit:"contain"},null,8,["src","preview-src-list"])):re("",!0)]),_:1}),e(u,{label:"\u82F1\u96C4\u540D\u79F0",prop:"name","min-width":"160","show-tooltip-when-overflow":""}),e(u,{label:"\u82F1\u96C4\u5B9A\u4F4D",prop:"types","min-width":"60"},{default:o(({row:t})=>{var y;return[r(L((y=a(_).find(j=>j.id===t.types))==null?void 0:y.name),1)]}),_:1}),e(u,{label:"\u64CD\u4F5C",width:"120",fixed:"right"},{default:o(({row:t})=>[g((i(),m(p,{type:"primary",link:""},{default:o(()=>[e(E,{to:{path:a(k)("hero:add/edit"),query:{id:t.id}}},{default:o(()=>[r(" \u7F16\u8F91 ")]),_:2},1032,["to"])]),_:2},1024)),[[b,["hero:edit","hero:add/edit"]]]),g((i(),m(p,{type:"danger",link:"",onClick:y=>T(t.id)},{default:o(()=>[r(" \u5220\u9664 ")]),_:2},1032,["onClick"])),[[b,["hero:del"]]])]),_:1})]),_:1},8,["data"])),[[S,a(d).loading]]),B("div",_e,[e($,{modelValue:a(d),"onUpdate:modelValue":s[3]||(s[3]=t=>ue(d)?d.value=t:null),onChange:a(c)},null,8,["modelValue","onChange"])])]),_:1})])}}});export{Ge as default};
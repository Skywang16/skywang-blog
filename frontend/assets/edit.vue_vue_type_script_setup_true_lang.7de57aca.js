import{C as k,D as C,v as E,F as V}from"./element-plus.8c1ec78f.js";import{a as x,c as y,b as R}from"./role.13aa6c64.js";import{P as h}from"./index.345a4530.js";import{f as B}from"./index.a935f15b.js";import{d as g,s as c,r as I,e as P,a1 as U,o as N,c as S,X as t,P as n,u as l}from"./@vue.a137a740.js";const j={class:"edit-popup"},K=g({__name:"edit",emits:["success","close"],setup(q,{expose:f,emit:p}){const i=c(),r=c(),m=I("add"),_=P(()=>m.value=="edit"?"\u7F16\u8F91\u89D2\u8272":"\u65B0\u589E\u89D2\u8272"),o=U({id:"",name:"",remark:"",sort:0,isDisable:0,menus:[]}),F={name:[{required:!0,message:"\u8BF7\u8F93\u5165\u540D\u79F0",trigger:["blur"]}]},b=async()=>{var e,a;await((e=i.value)==null?void 0:e.validate());const s={...o,menuIds:o.menus.join()};m.value=="edit"?await x(s):await y(s),(a=r.value)==null||a.close(),B.msgSuccess("\u64CD\u4F5C\u6210\u529F"),p("success")},w=()=>{p("close")};return f({open:(s="add")=>{var e;m.value=s,(e=r.value)==null||e.open()},setFormData:async s=>{const e=await R({id:s.id});for(const a in o)e[a]!=null&&e[a]!=null&&(o[a]=e[a])}}),(s,e)=>{const a=k,d=C,D=E,v=V;return N(),S("div",j,[t(h,{ref_key:"popupRef",ref:r,title:l(_),async:!0,width:"550px",onConfirm:b,onClose:w},{default:n(()=>[t(v,{class:"ls-form",ref_key:"formRef",ref:i,rules:F,model:l(o),"label-width":"60px"},{default:n(()=>[t(d,{label:"\u540D\u79F0",prop:"name"},{default:n(()=>[t(a,{class:"ls-input",modelValue:l(o).name,"onUpdate:modelValue":e[0]||(e[0]=u=>l(o).name=u),placeholder:"\u8BF7\u8F93\u5165\u540D\u79F0",clearable:""},null,8,["modelValue"])]),_:1}),t(d,{label:"\u5907\u6CE8",prop:"remark"},{default:n(()=>[t(a,{modelValue:l(o).remark,"onUpdate:modelValue":e[1]||(e[1]=u=>l(o).remark=u),type:"textarea",autosize:{minRows:4,maxRows:6},placeholder:"\u8BF7\u8F93\u5165\u5907\u6CE8",maxlength:"200","show-word-limit":""},null,8,["modelValue"])]),_:1}),t(d,{label:"\u6392\u5E8F",prop:"sort"},{default:n(()=>[t(D,{modelValue:l(o).sort,"onUpdate:modelValue":e[2]||(e[2]=u=>l(o).sort=u)},null,8,["modelValue"])]),_:1})]),_:1},8,["model"])]),_:1},8,["title"])])}}});export{K as _};
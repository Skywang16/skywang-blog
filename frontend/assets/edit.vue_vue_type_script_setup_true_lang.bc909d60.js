import{C as k,D as B,v as D,F as E}from"./element-plus.8c1ec78f.js";import{u as V,v as x,x as y}from"./information.8c78de9c.js";import{P as R}from"./index.345a4530.js";import{f as h}from"./index.a935f15b.js";import{d as g,s as c,r as S,e as I,a1 as P,o as U,c as N,X as s,P as n,u as t}from"./@vue.a137a740.js";const j={class:"edit-popup"},K=g({__name:"edit",emits:["success","close"],setup(q,{expose:f,emit:d}){const i=c(),r=c(),m=S("add"),_=I(()=>m.value=="edit"?"\u7F16\u8F91\u6B66\u5668\u7CFB\u5217":"\u65B0\u589E\u6B66\u5668\u7CFB\u5217"),o=P({id:"",name:"",remark:"",sort:0,isDisable:0,menus:[]}),F={name:[{required:!0,message:"\u8BF7\u8F93\u5165\u540D\u79F0",trigger:["blur"]}]},w=async()=>{var e,a;await((e=i.value)==null?void 0:e.validate());const u={...o,menuIds:o.menus.join()};m.value=="edit"?await V(u):await x(u),(a=r.value)==null||a.close(),h.msgSuccess("\u64CD\u4F5C\u6210\u529F"),d("success")},v=()=>{d("close")};return f({open:(u="add")=>{var e;m.value=u,(e=r.value)==null||e.open()},setFormData:async u=>{const e=await y({id:u.id});for(const a in o)e[a]!=null&&e[a]!=null&&(o[a]=e[a])}}),(u,e)=>{const a=k,p=B,C=D,b=E;return U(),N("div",j,[s(R,{ref_key:"popupRef",ref:r,title:t(_),async:!0,width:"550px",onConfirm:w,onClose:v},{default:n(()=>[s(b,{class:"ls-form",ref_key:"formRef",ref:i,rules:F,model:t(o),"label-width":"60px"},{default:n(()=>[s(p,{label:"\u540D\u79F0",prop:"name"},{default:n(()=>[s(a,{class:"ls-input",modelValue:t(o).name,"onUpdate:modelValue":e[0]||(e[0]=l=>t(o).name=l),placeholder:"\u8BF7\u8F93\u5165\u540D\u79F0",clearable:""},null,8,["modelValue"])]),_:1}),s(p,{label:"\u5907\u6CE8",prop:"remark"},{default:n(()=>[s(a,{modelValue:t(o).remark,"onUpdate:modelValue":e[1]||(e[1]=l=>t(o).remark=l),type:"textarea",autosize:{minRows:4,maxRows:6},placeholder:"\u8BF7\u8F93\u5165\u5907\u6CE8",maxlength:"200","show-word-limit":""},null,8,["modelValue"])]),_:1}),s(p,{label:"\u6392\u5E8F",prop:"sort"},{default:n(()=>[s(C,{modelValue:t(o).sort,"onUpdate:modelValue":e[2]||(e[2]=l=>t(o).sort=l)},null,8,["modelValue"])]),_:1})]),_:1},8,["model"])]),_:1},8,["title"])])}}});export{K as _};
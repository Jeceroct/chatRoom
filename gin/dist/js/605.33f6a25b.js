"use strict";(self["webpackChunkchatroom"]=self["webpackChunkchatroom"]||[]).push([[605],{1241:function(e,t){t.A=(e,t)=>{const s=e.__vccOpts||e;for(const[a,n]of t)s[a]=n;return s}},3531:function(e,t,s){s.d(t,{A:function(){return i}});var a=s(6768);const n={id:"routeMask"};function l(e,t,s,l,u,o){return(0,a.uX)(),(0,a.CE)("div",n," 加载中... ")}var u={name:"routeMask",data(){return{routeMaskEle:null}},mounted(){this.routeMaskEle=document.querySelector("#routeMask"),setTimeout((()=>{this.routeMaskEle.classList.add("enter"),setTimeout((()=>{this.routeMaskEle.classList.add("waiting"),this.routeMaskEle.classList.remove("enter")}),300)}),100)}},o=s(1241);const r=(0,o.A)(u,[["render",l]]);var i=r},605:function(e,t,s){s.r(t),s.d(t,{default:function(){return f}});s(4114),s(8992),s(3949);var a=s(6768),n=s(5130),l=s(3531),u=s(144),o=s(9294),r=s(1387),i=s(2251),c=s(1219);const d={class:"container"},v={class:"inputBox id"},p={class:"inputBox password"},m={class:"inputBox name"},L={class:"inputBox title"};var k={__name:"signup",setup(e){const t=(0,r.rd)(),s=(0,u.KR)(""),k=(0,u.KR)(""),h=(0,u.KR)(""),f=(0,u.KR)(""),E=(0,u.KR)(!1),y=()=>{s.value="";const e=document.querySelector("#routeMask");e.classList.add("leave"),e.classList.remove("waiting"),setTimeout((()=>{t.push("/login")}),200)},S=()=>{if(h.value=h.value.trim(),k.value=k.value.trim(),s.value=s.value.trim(),f.value=f.value.trim(),""===h.value){const e=document.querySelector(".name");e.classList.remove("hasValue"),e.classList.add("error"),e.querySelector("input").addEventListener("input",(()=>{e.classList.remove("error")}))}if(""===k.value){const e=document.querySelector(".password");e.classList.remove("hasValue"),e.classList.add("error"),e.querySelector("input").addEventListener("input",(()=>{e.classList.remove("error")}))}if(""===s.value){const e=document.querySelector(".id");e.querySelector("span").innerHTML="此项不能为空",e.classList.remove("hasValue"),e.classList.add("error"),e.querySelector("input").addEventListener("input",(()=>{e.classList.remove("error")}))}if(""===f.value){const e=document.querySelector(".title");e.classList.remove("hasValue")}if(""===h.value||""===k.value||""===s.value||E.value)return;const e=new FormData;e.append("id",s.value),e.append("password",k.value),e.append("name",h.value),e.append("title",f.value),o.A.post("/signup",e,{}).then((e=>{if(console.log(e),200==e.code){const e=document.querySelector("#routeMask");e.classList.add("leave"),e.classList.remove("waiting"),c.nk.success("注册成功"),setTimeout((()=>{t.push("/login")}),200)}}))};return(0,a.sV)((()=>{(0,i.B)();const e=document.querySelectorAll(".input");e.forEach((e=>{e.addEventListener("input",(()=>{""!==e.value?e.parentElement.classList.add("hasValue"):e.parentElement.classList.remove("hasValue")}))}));const t=document.querySelector(".id .input"),a=document.querySelector(".id span");t.addEventListener("input",(()=>{o.A.post("/checkIdUsed",{id:s.value},{}).then((e=>{console.log(e),"501"===e.code?(a.innerHTML="此id已被使用",t.parentElement.classList.remove("hasValue"),t.parentElement.classList.add("error"),E.value=!0):(a.innerHTML="",t.parentElement.classList.remove("error"),E.value=!1)}))}))})),(e,t)=>{const u=(0,a.g2)("el-form");return(0,a.uX)(),(0,a.CE)(a.FK,null,[(0,a.bF)(l.A),(0,a.Lk)("div",d,[t[8]||(t[8]=(0,a.Lk)("h1",null,"作为此聊天室的新用户进入",-1)),(0,a.bF)(u,{onSubmit:(0,n.D$)(S,["prevent"])},{default:(0,a.k6)((()=>[(0,a.Lk)("div",v,[(0,a.bo)((0,a.Lk)("input",{"onUpdate:modelValue":t[0]||(t[0]=e=>s.value=e),class:"input",placeholder:"取一个独特的id"},null,512),[[n.Jo,s.value]]),t[4]||(t[4]=(0,a.Lk)("span",null,"此项不能为空",-1))]),(0,a.Lk)("div",p,[(0,a.bo)((0,a.Lk)("input",{"onUpdate:modelValue":t[1]||(t[1]=e=>k.value=e),class:"input",placeholder:"取一个安全的密码"},null,512),[[n.Jo,k.value]]),t[5]||(t[5]=(0,a.Lk)("span",null,"此项不能为空",-1))]),(0,a.Lk)("div",m,[(0,a.bo)((0,a.Lk)("input",{"onUpdate:modelValue":t[2]||(t[2]=e=>h.value=e),class:"input",placeholder:"输入你的昵称"},null,512),[[n.Jo,h.value]]),t[6]||(t[6]=(0,a.Lk)("span",null,"此项不能为空",-1))]),(0,a.Lk)("div",L,[(0,a.bo)((0,a.Lk)("input",{"onUpdate:modelValue":t[3]||(t[3]=e=>f.value=e),class:"input",placeholder:"给自己一个头衔吧，不过有没有无所谓了"},null,512),[[n.Jo,f.value]])]),(0,a.Lk)("button",{id:"loginBtn",class:"loginBtn",onClick:y},"去登录"),t[7]||(t[7]=(0,a.Lk)("button",{id:"submitBtn",class:"submitBtn",type:"submit"},"注册",-1))])),_:1})])],64)}}};const h=k;var f=h}}]);
//# sourceMappingURL=605.33f6a25b.js.map
(function(){"use strict";var e={9294:function(e,n,t){t(4114);var r=t(4373),o=t(1219),i=t(9325);const u=r.A.create({baseURL:""});var a=!1;u.interceptors.request.use((e=>{e.headers["Content-Type"]="application/json;charset=utf-8";let n=JSON.parse(localStorage.getItem("xm-user")||"{}");return e.headers["token"]=n.token||"",e}),(e=>(o.nk.error("无法连接至服务器"),Promise.reject(e)))),u.interceptors.response.use((e=>{let n=e.data;return"blob"===e.config.responseType||"401"===n.code&&(o.nk.error(n.msg),i.A.push("/login")),n}),(e=>(a||(a=!0,404===e.response.status||(500===e.response.status?o.nk.error("无法连接至服务器"):501===e.response.status?o.nk.error("文件接收失败"):(o.nk.error("无法连接至服务器"),console.error(e.message))),setTimeout((()=>{a=!1}),4e3),console.warn(e.response)),Promise.reject(e)))),n.A=u},6982:function(e,n,t){var r=t(5130),o=t(6768),i=t(2251),u=t(1387),a={__name:"App",setup(e){const n=(0,u.rd)(),t=e=>new Promise((n=>setTimeout(n,e)));return(0,o.sV)((async()=>{for(;;)await t(1e3),(0,i.B)(n)})),(e,n)=>{const t=(0,o.g2)("router-view");return(0,o.uX)(),(0,o.Wv)(t)}}};const c=a;var s=c,f=t(9325),l=t(5835),d=(t(4188),t(7477)),p=t(8349),h=t(8950),m=t(292),v=t(2353),g=t(92),b=t(4996);h.Yv.add(v.X7I,g.C91,b.Cvc);const y=(0,r.Ef)(s);for(const[k,w]of Object.entries(d))y.component(k,w);y.use(l.A),y.component("font-awesome-icon",m.gc),y.component("EmojiPicker",p.A),y.use(f.A),y.mount("#app")},9325:function(e,n,t){var r=t(1387);const o=[{path:"/address",name:"address",component:()=>t.e(567).then(t.bind(t,3567))},{path:"/signup",name:"signup",component:()=>t.e(605).then(t.bind(t,605))},{path:"/login",name:"login",component:()=>t.e(834).then(t.bind(t,3834))},{path:"/room",name:"room",component:()=>t.e(168).then(t.bind(t,1168))},{path:"/",name:"starting",component:()=>t.e(75).then(t.bind(t,1075))}],i=(0,r.aE)({history:(0,r.Bt)(),routes:o});n.A=i},2251:function(e,n,t){t.d(n,{B:function(){return i}});t(4114);var r=t(9294),o="";function i(e){r.A.post("/getStatus").then((n=>{switch(n.status){case"checkRoom":if("checkRoom"==o)return;console.log(n),o="checkRoom",e.push("/address");break;case"checkUser":if("checkUser"==o)return;console.log(n),o="checkUser",e.push("/login");break;case"start":if("start"==o)return;console.log(n),o="start",e.push("/room");break}}))}}},n={};function t(r){var o=n[r];if(void 0!==o)return o.exports;var i=n[r]={exports:{}};return e[r].call(i.exports,i,i.exports,t),i.exports}t.m=e,function(){var e=[];t.O=function(n,r,o,i){if(!r){var u=1/0;for(f=0;f<e.length;f++){r=e[f][0],o=e[f][1],i=e[f][2];for(var a=!0,c=0;c<r.length;c++)(!1&i||u>=i)&&Object.keys(t.O).every((function(e){return t.O[e](r[c])}))?r.splice(c--,1):(a=!1,i<u&&(u=i));if(a){e.splice(f--,1);var s=o();void 0!==s&&(n=s)}}return n}i=i||0;for(var f=e.length;f>0&&e[f-1][2]>i;f--)e[f]=e[f-1];e[f]=[r,o,i]}}(),function(){t.n=function(e){var n=e&&e.__esModule?function(){return e["default"]}:function(){return e};return t.d(n,{a:n}),n}}(),function(){t.d=function(e,n){for(var r in n)t.o(n,r)&&!t.o(e,r)&&Object.defineProperty(e,r,{enumerable:!0,get:n[r]})}}(),function(){t.f={},t.e=function(e){return Promise.all(Object.keys(t.f).reduce((function(n,r){return t.f[r](e,n),n}),[]))}}(),function(){t.u=function(e){return"js/"+e+"."+{75:"51254ad4",168:"956cd3f9",567:"7b6590f9",605:"33f6a25b",834:"a85bfd40"}[e]+".js"}}(),function(){t.miniCssF=function(e){return"css/"+e+"."+{75:"6080a70f",168:"579109f7",567:"42d5e8dc",605:"42d5e8dc",834:"42d5e8dc"}[e]+".css"}}(),function(){t.g=function(){if("object"===typeof globalThis)return globalThis;try{return this||new Function("return this")()}catch(e){if("object"===typeof window)return window}}()}(),function(){t.o=function(e,n){return Object.prototype.hasOwnProperty.call(e,n)}}(),function(){var e={},n="chatroom:";t.l=function(r,o,i,u){if(e[r])e[r].push(o);else{var a,c;if(void 0!==i)for(var s=document.getElementsByTagName("script"),f=0;f<s.length;f++){var l=s[f];if(l.getAttribute("src")==r||l.getAttribute("data-webpack")==n+i){a=l;break}}a||(c=!0,a=document.createElement("script"),a.charset="utf-8",a.timeout=120,t.nc&&a.setAttribute("nonce",t.nc),a.setAttribute("data-webpack",n+i),a.src=r),e[r]=[o];var d=function(n,t){a.onerror=a.onload=null,clearTimeout(p);var o=e[r];if(delete e[r],a.parentNode&&a.parentNode.removeChild(a),o&&o.forEach((function(e){return e(t)})),n)return n(t)},p=setTimeout(d.bind(null,void 0,{type:"timeout",target:a}),12e4);a.onerror=d.bind(null,a.onerror),a.onload=d.bind(null,a.onload),c&&document.head.appendChild(a)}}}(),function(){t.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})}}(),function(){t.p=""}(),function(){if("undefined"!==typeof document){var e=function(e,n,r,o,i){var u=document.createElement("link");u.rel="stylesheet",u.type="text/css",t.nc&&(u.nonce=t.nc);var a=function(t){if(u.onerror=u.onload=null,"load"===t.type)o();else{var r=t&&t.type,a=t&&t.target&&t.target.href||n,c=new Error("Loading CSS chunk "+e+" failed.\n("+r+": "+a+")");c.name="ChunkLoadError",c.code="CSS_CHUNK_LOAD_FAILED",c.type=r,c.request=a,u.parentNode&&u.parentNode.removeChild(u),i(c)}};return u.onerror=u.onload=a,u.href=n,r?r.parentNode.insertBefore(u,r.nextSibling):document.head.appendChild(u),u},n=function(e,n){for(var t=document.getElementsByTagName("link"),r=0;r<t.length;r++){var o=t[r],i=o.getAttribute("data-href")||o.getAttribute("href");if("stylesheet"===o.rel&&(i===e||i===n))return o}var u=document.getElementsByTagName("style");for(r=0;r<u.length;r++){o=u[r],i=o.getAttribute("data-href");if(i===e||i===n)return o}},r=function(r){return new Promise((function(o,i){var u=t.miniCssF(r),a=t.p+u;if(n(u,a))return o();e(r,a,null,o,i)}))},o={524:0};t.f.miniCss=function(e,n){var t={75:1,168:1,567:1,605:1,834:1};o[e]?n.push(o[e]):0!==o[e]&&t[e]&&n.push(o[e]=r(e).then((function(){o[e]=0}),(function(n){throw delete o[e],n})))}}}(),function(){var e={524:0};t.f.j=function(n,r){var o=t.o(e,n)?e[n]:void 0;if(0!==o)if(o)r.push(o[2]);else{var i=new Promise((function(t,r){o=e[n]=[t,r]}));r.push(o[2]=i);var u=t.p+t.u(n),a=new Error,c=function(r){if(t.o(e,n)&&(o=e[n],0!==o&&(e[n]=void 0),o)){var i=r&&("load"===r.type?"missing":r.type),u=r&&r.target&&r.target.src;a.message="Loading chunk "+n+" failed.\n("+i+": "+u+")",a.name="ChunkLoadError",a.type=i,a.request=u,o[1](a)}};t.l(u,c,"chunk-"+n,n)}},t.O.j=function(n){return 0===e[n]};var n=function(n,r){var o,i,u=r[0],a=r[1],c=r[2],s=0;if(u.some((function(n){return 0!==e[n]}))){for(o in a)t.o(a,o)&&(t.m[o]=a[o]);if(c)var f=c(t)}for(n&&n(r);s<u.length;s++)i=u[s],t.o(e,i)&&e[i]&&e[i][0](),e[i]=0;return t.O(f)},r=self["webpackChunkchatroom"]=self["webpackChunkchatroom"]||[];r.forEach(n.bind(null,0)),r.push=n.bind(null,r.push.bind(r))}();var r=t.O(void 0,[504],(function(){return t(6982)}));r=t.O(r)})();
//# sourceMappingURL=app.4563564f.js.map
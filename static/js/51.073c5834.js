"use strict";(self["webpackChunkkube_manage_web"]=self["webpackChunkkube_manage_web"]||[]).push([[51],{9922:function(e,a,t){t.d(a,{FI:function(){return n},PY:function(){return l},Yw:function(){return i},mY:function(){return r}});var s=t(4471);function n(e){return(0,s.Z)({url:"/api/k8s/secret/list",method:"get",params:e})}function l(e){return(0,s.Z)({url:"/api/k8s/secret/detail",method:"get",params:e})}function i(e){return(0,s.Z)({url:"/api/k8s/secret/update",method:"put",params:e})}function r(e){return(0,s.Z)({url:"/api/k8s/secret/del",method:"delete",params:e})}},9051:function(e,a,t){t.r(a),t.d(a,{default:function(){return x}});var s=t(3396),n=t(7139);const l=e=>((0,s.dD)("data-v-2a91d83e"),e=e(),(0,s.Cn)(),e),i={class:"pvc"},r=l((()=>(0,s._)("span",null,"命名空间: ",-1))),c=(0,s.Uk)("刷新"),p=(0,s.Uk)("创建"),m=(0,s.Uk)("搜索"),o={class:"pvc-body-pvcname"},u=(0,s.Uk)("YAML"),d=(0,s.Uk)("删除"),g={class:"dialog-footer"},h=(0,s.Uk)("取 消"),v=(0,s.Uk)("更 新");function f(e,a,t,l,f,w){const _=(0,s.up)("el-option"),b=(0,s.up)("el-select"),y=(0,s.up)("el-col"),D=(0,s.up)("el-button"),P=(0,s.up)("el-row"),C=(0,s.up)("el-card"),W=(0,s.up)("el-input"),k=(0,s.up)("el-table-column"),z=(0,s.up)("el-tag"),V=(0,s.up)("el-popover"),L=(0,s.up)("el-table"),x=(0,s.up)("el-pagination"),Y=(0,s.up)("codemirror"),S=(0,s.up)("el-dialog"),U=(0,s.Q2)("loading");return(0,s.wg)(),(0,s.iD)("div",i,[(0,s.Wm)(P,null,{default:(0,s.w5)((()=>[(0,s.Wm)(y,{span:24},{default:(0,s.w5)((()=>[(0,s._)("div",null,[(0,s.Wm)(C,{class:"pvc-head-card",shadow:"never","body-style":{padding:"10px"}},{default:(0,s.w5)((()=>[(0,s.Wm)(P,null,{default:(0,s.w5)((()=>[(0,s.Wm)(y,{span:6},{default:(0,s.w5)((()=>[(0,s._)("div",null,[r,(0,s.Wm)(b,{modelValue:f.namespaceValue,"onUpdate:modelValue":a[0]||(a[0]=e=>f.namespaceValue=e),filterable:"",placeholder:"请选择"},{default:(0,s.w5)((()=>[((0,s.wg)(!0),(0,s.iD)(s.HY,null,(0,s.Ko)(f.namespaceList,((e,a)=>((0,s.wg)(),(0,s.j4)(_,{key:a,label:e.metadata.name,value:e.metadata.name},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])])),_:1}),(0,s.Wm)(y,{span:2,offset:16},{default:(0,s.w5)((()=>[(0,s._)("div",null,[(0,s.Wm)(D,{style:{"border-radius":"2px"},icon:"Refresh",plain:"",onClick:a[1]||(a[1]=e=>w.getPvcs())},{default:(0,s.w5)((()=>[c])),_:1})])])),_:1})])),_:1})])),_:1})])])),_:1}),(0,s.Wm)(y,{span:24},{default:(0,s.w5)((()=>[(0,s._)("div",null,[(0,s.Wm)(C,{class:"pvc-head-card",shadow:"never","body-style":{padding:"10px"}},{default:(0,s.w5)((()=>[(0,s.Wm)(P,null,{default:(0,s.w5)((()=>[(0,s.Wm)(y,{span:2},{default:(0,s.w5)((()=>[(0,s._)("div",null,[(0,s.Wm)(D,{disabled:"",style:{"border-radius":"2px"},icon:"Edit",type:"primary"},{default:(0,s.w5)((()=>[p])),_:1})])])),_:1}),(0,s.Wm)(y,{span:6},{default:(0,s.w5)((()=>[(0,s._)("div",null,[(0,s.Wm)(W,{class:"pvc-head-search",clearable:"",placeholder:"请输入",modelValue:f.searchInput,"onUpdate:modelValue":a[2]||(a[2]=e=>f.searchInput=e)},null,8,["modelValue"]),(0,s.Wm)(D,{style:{"border-radius":"2px"},icon:"Search",type:"primary",plain:"",onClick:a[3]||(a[3]=e=>w.getPvcs())},{default:(0,s.w5)((()=>[m])),_:1})])])),_:1})])),_:1})])),_:1})])])),_:1}),(0,s.Wm)(y,{span:24},{default:(0,s.w5)((()=>[(0,s._)("div",null,[(0,s.Wm)(C,{class:"pvc-body-card",shadow:"never","body-style":{padding:"5px"}},{default:(0,s.w5)((()=>[(0,s.wy)(((0,s.wg)(),(0,s.j4)(L,{style:{width:"100%","font-size":"12px","margin-bottom":"10px"},data:f.pvcList},{default:(0,s.w5)((()=>[(0,s.Wm)(k,{width:"20"}),(0,s.Wm)(k,{align:"left",label:"PVC名"},{default:(0,s.w5)((e=>[(0,s._)("a",o,(0,n.zw)(e.row.metadata.name),1)])),_:1}),(0,s.Wm)(k,{align:"center",label:"标签"},{default:(0,s.w5)((e=>[((0,s.wg)(!0),(0,s.iD)(s.HY,null,(0,s.Ko)(e.row.metadata.labels,((e,a)=>((0,s.wg)(),(0,s.iD)("div",{key:a},[(0,s.Wm)(V,{placement:"right",width:200,trigger:"hover",content:a+":"+e},{reference:(0,s.w5)((()=>[(0,s.Wm)(z,{style:{"margin-bottom":"5px"},type:"warning"},{default:(0,s.w5)((()=>[(0,s.Uk)((0,n.zw)(w.ellipsis(a+":"+e)),1)])),_:2},1024)])),_:2},1032,["content"])])))),128))])),_:1}),(0,s.Wm)(k,{align:"center",label:"状态"},{default:(0,s.w5)((e=>[(0,s._)("span",{class:(0,n.C_)(["Bound"===e.row.status.phase?"success-status":"error-status"])},(0,n.zw)(e.row.status.phase),3)])),_:1}),(0,s.Wm)(k,{align:"center",prop:"status.capacity.storage",label:"容量"}),(0,s.Wm)(k,{align:"center",prop:"status.accessModes[0]",label:"访问模式"}),(0,s.Wm)(k,{align:"center",prop:"spec.storageClassName",label:"StorageClass"}),(0,s.Wm)(k,{align:"center","min-width":"100",label:"创建时间"},{default:(0,s.w5)((e=>[(0,s.Wm)(z,{type:"info"},{default:(0,s.w5)((()=>[(0,s.Uk)((0,n.zw)(w.timeTrans(e.row.metadata.creationTimestamp)),1)])),_:2},1024)])),_:1}),(0,s.Wm)(k,{align:"center",label:"操作",width:"200"},{default:(0,s.w5)((e=>[(0,s.Wm)(D,{size:"small",style:{"border-radius":"2px"},icon:"Edit",type:"primary",plain:"",onClick:a=>w.getPvcDetail(e)},{default:(0,s.w5)((()=>[u])),_:2},1032,["onClick"]),(0,s.Wm)(D,{size:"small",style:{"border-radius":"2px"},icon:"Delete",type:"danger",onClick:a=>w.handleConfirm(e,"删除",w.delPvc)},{default:(0,s.w5)((()=>[d])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"])),[[U,f.appLoading]]),(0,s.Wm)(x,{class:"pvc-body-pagination",background:"",onSizeChange:w.handleSizeChange,onCurrentChange:w.handleCurrentChange,"current-page":f.currentPage,"page-sizes":f.pagesizeList,"page-size":f.pagesize,layout:"total, sizes, prev, pager, next, jumper",total:f.pvcTotal},null,8,["onSizeChange","onCurrentChange","current-page","page-sizes","page-size","total"])])),_:1})])])),_:1})])),_:1}),(0,s.Wm)(S,{title:"YAML信息",modelValue:f.yamlDialog,"onUpdate:modelValue":a[6]||(a[6]=e=>f.yamlDialog=e),width:"45%",top:"5%"},{footer:(0,s.w5)((()=>[(0,s._)("span",g,[(0,s.Wm)(D,{onClick:a[4]||(a[4]=e=>f.yamlDialog=!1)},{default:(0,s.w5)((()=>[h])),_:1}),(0,s.Wm)(D,{type:"primary",onClick:a[5]||(a[5]=e=>w.updatePvc())},{default:(0,s.w5)((()=>[v])),_:1})])])),default:(0,s.w5)((()=>[(0,s.Wm)(Y,{value:f.contentYaml,border:"",options:f.cmOptions,height:"500",style:{"font-size":"14px"},onChange:w.onChange},null,8,["value","options","onChange"])])),_:1},8,["modelValue"])])}var w=t(7973),_=t(5959),b=t(1391),y=t(3794),D=t(4471);function P(e){return(0,D.Z)({url:"/api/k8s/persistentvolumeclaim/list",method:"get",params:e})}function C(e){return(0,D.Z)({url:"/api/k8s/persistentvolumeclaim/detail",method:"get",params:e})}function W(e){return(0,D.Z)({url:"/api/k8s/persistentvolumeclaim/update",method:"put",params:e})}var k=t(9922),z={data(){return{cmOptions:w.Z.cmOptions,contentYaml:"",currentPage:1,pagesize:10,pagesizeList:[10,20,30],searchInput:"",namespaceValue:"default",namespaceList:[],namespaceListUrl:w.Z.k8sNamespaceList,appLoading:!1,pvcList:[],pvcTotal:0,getPvcsData:{params:{filter_name:"",namespace:"",page:"",limit:""}},pvcDetail:{},getPvcDetailData:{params:{pvc_name:"",namespace:""}},yamlDialog:!1,updatePvcData:{params:{namespace:"",content:""}},delPvcData:{params:{pvc_name:"",namespace:""}}}},methods:{transYaml(e){return b.stringify(e)},transObj(e){return _.ZP.load(e)},onChange(e){this.contentYaml=e},handleSizeChange(e){this.pagesize=e,this.getPvcs()},handleCurrentChange(e){this.currentPage=e,this.getPvcs()},handleClose(e){this.$confirm("确认关闭？").then((()=>{e()})).catch((()=>{}))},ellipsis(e){return e.length>15?e.substring(0,15)+"...":e},timeTrans(e){let a=new Date(new Date(e).getTime()+288e5);return a=a.toJSON(),a=a.substring(0,19).replace("T"," "),a},restartTotal(e){let a,t=0,s=e.row.status.containerStatuses;for(a in s)t+=s[a].restartCount;return t},getNamespaces(){(0,y.I1)().then((e=>{this.namespaceList=e.data.items})).catch((e=>{this.$message.error({message:e.msg})}))},getPvcs(){this.appLoading=!0,this.getPvcsData.params.filter_name=this.searchInput,this.getPvcsData.params.namespace=this.namespaceValue,this.getPvcsData.params.page=this.currentPage,this.getPvcsData.params.limit=this.pagesize,P(this.getPvcsData.params).then((e=>{this.pvcList=e.data.items,this.pvcTotal=e.data.total})).catch((e=>{this.$message.error({message:e.msg})})),this.appLoading=!1},getPvcDetail(e){this.getPvcDetailData.params.pvc_name=e.row.metadata.name,this.getPvcDetailData.params.namespace=this.namespaceValue,C(this.getPvcDetailData.params).then((e=>{this.pvcDetail=e.data,this.contentYaml=this.transYaml(this.pvcDetail),this.yamlDialog=!0})).catch((e=>{this.$message.error({message:e.msg})}))},updatePvc(){let e=JSON.stringify(this.transObj(this.contentYaml));this.updatePvcData.params.namespace=this.namespaceValue,this.updatePvcData.params.content=e,W(this.updatePvcData.params).then((e=>{this.$message.success({message:e.msg})})).catch((e=>{this.$message.error({message:e.msg})})),this.yamlDialog=!1},delPvc(e){this.delPvcData.params.pvc_name=e.row.metadata.name,this.delPvcData.params.namespace=this.namespaceValue,(0,k.mY)(this.delPvcData.params).then((e=>{this.getPvcs(),this.$message.success({message:e.msg})})).catch((e=>{this.$message.error({message:e.msg})}))},handleConfirm(e,a,t){this.confirmContent="确认继续 "+a+" 操作吗？",this.$confirm(this.confirmContent,"提示",{confirmButtonText:"确定",cancelButtonText:"取消"}).then((()=>{t(e)})).catch((()=>{this.$message.info({message:"已取消操作"})}))}},watch:{namespaceValue:{handler(){localStorage.setItem("namespace",this.namespaceValue),this.currentPage=1,this.getPvcs()}}},beforeMount(){void 0!==localStorage.getItem("namespace")&&null!==localStorage.getItem("namespace")&&(this.namespaceValue=localStorage.getItem("namespace")),this.getNamespaces(),this.getPvcs()}},V=t(89);const L=(0,V.Z)(z,[["render",f],["__scopeId","data-v-2a91d83e"]]);var x=L}}]);
//# sourceMappingURL=51.073c5834.js.map
"use strict";(self["webpackChunkkube_manage_web"]=self["webpackChunkkube_manage_web"]||[]).push([[941],{1941:function(e,a,t){t.r(a),t.d(a,{default:function(){return z}});var s=t(3396),n=t(7139);const l={class:"pv"},i=(0,s.Uk)("创建"),r=(0,s.Uk)("搜索"),p=(0,s.Uk)("刷新"),m={class:"pv-body-pvname"},o=(0,s.Uk)("YAML"),u=(0,s.Uk)("删除"),d={class:"dialog-footer"},c=(0,s.Uk)("取 消"),g=(0,s.Uk)("更 新");function h(e,a,t,h,v,f){const w=(0,s.up)("el-button"),P=(0,s.up)("el-col"),_=(0,s.up)("el-input"),D=(0,s.up)("el-row"),b=(0,s.up)("el-card"),y=(0,s.up)("el-table-column"),C=(0,s.up)("el-tag"),k=(0,s.up)("el-table"),W=(0,s.up)("el-pagination"),z=(0,s.up)("codemirror"),L=(0,s.up)("el-dialog"),x=(0,s.Q2)("loading");return(0,s.wg)(),(0,s.iD)("div",l,[(0,s.Wm)(D,null,{default:(0,s.w5)((()=>[(0,s.Wm)(P,{span:24},{default:(0,s.w5)((()=>[(0,s._)("div",null,[(0,s.Wm)(b,{class:"pv-head-card",shadow:"never","body-style":{padding:"10px"}},{default:(0,s.w5)((()=>[(0,s.Wm)(D,null,{default:(0,s.w5)((()=>[(0,s.Wm)(P,{span:2},{default:(0,s.w5)((()=>[(0,s._)("div",null,[(0,s.Wm)(w,{disabled:"",style:{"border-radius":"2px"},icon:"Edit",type:"primary"},{default:(0,s.w5)((()=>[i])),_:1})])])),_:1}),(0,s.Wm)(P,{span:6},{default:(0,s.w5)((()=>[(0,s._)("div",null,[(0,s.Wm)(_,{class:"pv-head-search",clearable:"",placeholder:"请输入",modelValue:v.searchInput,"onUpdate:modelValue":a[0]||(a[0]=e=>v.searchInput=e)},null,8,["modelValue"]),(0,s.Wm)(w,{style:{"border-radius":"2px"},icon:"Search",type:"primary",plain:"",onClick:a[1]||(a[1]=e=>f.getPvs())},{default:(0,s.w5)((()=>[r])),_:1})])])),_:1}),(0,s.Wm)(P,{span:2,offset:14},{default:(0,s.w5)((()=>[(0,s._)("div",null,[(0,s.Wm)(w,{style:{"border-radius":"2px"},icon:"Refresh",plain:"",onClick:a[2]||(a[2]=e=>f.getPvs())},{default:(0,s.w5)((()=>[p])),_:1})])])),_:1})])),_:1})])),_:1})])])),_:1}),(0,s.Wm)(P,{span:24},{default:(0,s.w5)((()=>[(0,s._)("div",null,[(0,s.Wm)(b,{class:"pv-body-card",shadow:"never","body-style":{padding:"5px"}},{default:(0,s.w5)((()=>[(0,s.wy)(((0,s.wg)(),(0,s.j4)(k,{style:{width:"100%","font-size":"12px","margin-bottom":"10px"},data:v.pvList},{default:(0,s.w5)((()=>[(0,s.Wm)(y,{width:"20"}),(0,s.Wm)(y,{align:"left",label:"PV名"},{default:(0,s.w5)((e=>[(0,s._)("a",m,(0,n.zw)(e.row.metadata.name),1)])),_:1}),(0,s.Wm)(y,{align:"center",label:"状态"},{default:(0,s.w5)((e=>[(0,s._)("span",{class:(0,n.C_)(["Bound"===e.row.status.phase?"success-status":"error-status"])},(0,n.zw)(e.row.status.phase),3)])),_:1}),(0,s.Wm)(y,{align:"center",prop:"spec.accessModes[0]",label:"访问模式"}),(0,s.Wm)(y,{align:"center",prop:"spec.capacity.storage",label:"容量"}),(0,s.Wm)(y,{align:"center",prop:"spec.claimRef.name",label:"PVC"}),(0,s.Wm)(y,{align:"center","min-width":"100",label:"创建时间"},{default:(0,s.w5)((e=>[(0,s.Wm)(C,{type:"info"},{default:(0,s.w5)((()=>[(0,s.Uk)((0,n.zw)(f.timeTrans(e.row.metadata.creationTimestamp)),1)])),_:2},1024)])),_:1}),(0,s.Wm)(y,{align:"center",label:"操作","min-width":"120"},{default:(0,s.w5)((e=>[(0,s.Wm)(w,{size:"small",style:{"border-radius":"2px"},icon:"Edit",type:"primary",plain:"",onClick:a=>f.getPvDetail(e)},{default:(0,s.w5)((()=>[o])),_:2},1032,["onClick"]),(0,s.Wm)(w,{size:"small",style:{"border-radius":"2px"},icon:"Delete",type:"danger",onClick:a=>f.handleConfirm(e,"删除",f.delPv)},{default:(0,s.w5)((()=>[u])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"])),[[x,v.appLoading]]),(0,s.Wm)(W,{class:"pv-body-pagination",background:"",onSizeChange:f.handleSizeChange,onCurrentChange:f.handleCurrentChange,"current-page":v.currentPage,"page-sizes":v.pagesizeList,"page-size":v.pagesize,layout:"total, sizes, prev, pager, next, jumper",total:v.pvTotal},null,8,["onSizeChange","onCurrentChange","current-page","page-sizes","page-size","total"])])),_:1})])])),_:1})])),_:1}),(0,s.Wm)(L,{title:"YAML信息",modelValue:v.yamlDialog,"onUpdate:modelValue":a[5]||(a[5]=e=>v.yamlDialog=e),width:"45%",top:"5%"},{footer:(0,s.w5)((()=>[(0,s._)("span",d,[(0,s.Wm)(w,{onClick:a[3]||(a[3]=e=>v.yamlDialog=!1)},{default:(0,s.w5)((()=>[c])),_:1}),(0,s.Wm)(w,{disabled:"",type:"primary",onClick:a[4]||(a[4]=e=>f.updatePv())},{default:(0,s.w5)((()=>[g])),_:1})])])),default:(0,s.w5)((()=>[(0,s.Wm)(z,{value:v.contentYaml,border:"",options:v.cmOptions,height:"500",style:{"font-size":"14px"},onChange:f.onChange},null,8,["value","options","onChange"])])),_:1},8,["modelValue"])])}var v=t(7973),f=t(5959),w=t(1391),P=t(4471);function _(e){return(0,P.Z)({url:"/api/k8s/persistentvolume/list",method:"get",params:e})}function D(e){return(0,P.Z)({url:"/api/k8s/persistentvolume/detail",method:"get",params:e})}function b(e){return(0,P.Z)({url:"/api/k8s/persistentvolume/update",method:"put",params:e})}function y(e){return(0,P.Z)({url:"/api/k8s/persistentvolume/del",method:"delete",params:e})}var C={data(){return{cmOptions:v.Z.cmOptions,contentYaml:"",currentPage:1,pagesize:10,pagesizeList:[10,20,30],searchInput:"",namespaceValue:"default",namespaceList:[],namespaceListUrl:v.Z.k8sNamespaceList,appLoading:!1,pvList:[],pvTotal:0,getPvsData:{url:v.Z.k8sPvList,params:{filter_name:"",namespace:"",page:"",limit:""}},pvDetail:{},getPvDetailData:{url:v.Z.k8sPvDetail,params:{name:"",namespace:""}},yamlDialog:!1,updatePvData:{url:v.Z.k8sPvUpdate,params:{namespace:"",content:""}},delPvData:{url:v.Z.k8spvDel,params:{name:"",namespace:""}}}},methods:{transYaml(e){return w.stringify(e)},transObj(e){return f.ZP.load(e)},onChange(e){this.contentYaml=e},handleSizeChange(e){this.pagesize=e,this.getPvs()},handleCurrentChange(e){this.currentPage=e,this.getPvs()},handleClose(e){this.$confirm("确认关闭？").then((()=>{e()})).catch((()=>{}))},ellipsis(e){return e.length>15?e.substring(0,15)+"...":e},timeTrans(e){let a=new Date(new Date(e).getTime()+288e5);return a=a.toJSON(),a=a.substring(0,19).replace("T"," "),a},specTrans(e){if(-1==e.indexOf("Ki"))return e;let a=e.slice(0,-2)/1024/1024;return a.toFixed(0)},getNamespaces(){getNamespecelist().then((e=>{this.namespaceList=e.data.items})).catch((e=>{this.$message.error({message:e.msg})}))},getPvs(){this.appLoading=!0,this.getPvsData.params.filter_name=this.searchInput,this.getPvsData.params.namespace=this.namespaceValue,this.getPvsData.params.page=this.currentPage,this.getPvsData.params.limit=this.pagesize,_(this.getPvsData.params).then((e=>{this.pvList=e.data.items,this.pvTotal=e.data.total})).catch((e=>{this.$message.error({message:e.msg})})),this.appLoading=!1},getPvDetail(e){this.getPvDetailData.params.name=e.row.metadata.name,this.getPvDetailData.params.namespace=this.namespaceValue,D(this.getPvDetailData.params).then((e=>{this.pvDetail=e.data,this.contentYaml=this.transYaml(this.pvDetail),this.yamlDialog=!0})).catch((e=>{this.$message.error({message:e.msg})}))},updatePv(){let e=JSON.stringify(this.transObj(this.contentYaml));this.updatePvData.params.namespace=this.namespaceValue,this.updatePvData.params.content=e,b(this.updatePvData.params).then((e=>{this.$message.success({message:e.msg})})).catch((e=>{this.$message.error({message:e.msg})})),this.yamlDialog=!1},delPv(e){this.delPvData.params.name=e.row.metadata.name,this.delPvData.params.namespace=this.namespaceValue,y(this.delPvData.params).then((e=>{this.getPvs(),this.$message.success({message:e.msg})})).catch((e=>{this.$message.error({message:e.msg})}))},handleConfirm(e,a,t){this.confirmContent="确认继续 "+a+" 操作吗？",this.$confirm(this.confirmContent,"提示",{confirmButtonText:"确定",cancelButtonText:"取消"}).then((()=>{t(e)})).catch((()=>{this.$message.info({message:"已取消操作"})}))}},beforeMount(){this.getPvs()}},k=t(89);const W=(0,k.Z)(C,[["render",h],["__scopeId","data-v-f36d00a0"]]);var z=W}}]);
//# sourceMappingURL=941.d010ac0b.js.map
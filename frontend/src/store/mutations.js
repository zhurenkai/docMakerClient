export default {
    addTab(state,api){
        state.selectedApi = api
        // 直接用indexOf(api)居然没用，应该是js重新给了内存地址
        let exits = state.apiTabs.filter(function(item,index){return item.id ==api.id })
        if(exits.length==0){
            state.apiTabs.push(api)
            state.activeApi = String(state.apiTabs.length-1)
        }else{
            let index = state.apiTabs.indexOf(exits[0])
            state.activeApi = String(index)

        }
    },
    removeTab(state,index){
        state.apiTabs.splice(index,1)
        if( index == state.activeApi ){
            let nextApi = state.apiTabs[parseInt(index)+1] || state.apiTabs[parseInt(index)-1]
            let nextIndex = state.apiTabs.indexOf(nextApi)
            state.activeApi = String(nextIndex)

        }else if(index < state.activeApi){

            state.activeApi = String(state.activeApi-1)
        }
    },
  userInfo (state,userInfo) {
        state.userInfo = userInfo
  }
}
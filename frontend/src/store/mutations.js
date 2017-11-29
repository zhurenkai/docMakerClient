export default {
    selectApi(state,api){
        state.selectedApi = api
        let index = state.apiTabs.indexOf(api)
        if(index==-1){
            state.apiTabs.push(api)
            state.activeApi = String(state.apiTabs.length-1)
        }else{
            state.activeApi = String(index)

        }
    },
    removeTab(state,index){
        state.apiTabs.splice(index,1)
        if( index == state.activeApi ){
            let nextApi = state.apiTabs[parseInt(index)+1] || state.apiTabs[parseInt(index)-1]
            let nextIndex = state.apiTabs.indexOf(nextApi)
            state.activeApi = String(nextIndex)
            console.log(state.activeApi)
        }else if(index < state.activeApi){
            console.log(index-1)
            state.activeApi = String(state.activeApi-1)
        }
    }
}
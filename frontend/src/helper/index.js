exports.install = function (Vue, options) {
    Vue.prototype.getApiData = function (response){
        if(response.data.code){
            this.$message.error(response.data.error)
            return false
        }else{
            if(arguments[1]){
                this.$message.success(arguments[1])
            }
            return response.data.data
        }
    };
};
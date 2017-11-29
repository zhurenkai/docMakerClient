exports.install = function (Vue, options) {
    Vue.prototype.getApiData = function (response){
        if(response.data.code){
            this.$message.error(response.data.error)
            return false
        }else{
            this.$message.success('成功')
            return response.data.data
        }
    };
};
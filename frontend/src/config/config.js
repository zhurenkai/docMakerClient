const proxy_domain = 'api/'
const prefix = 'api/'
export const getUri = (type,key ) =>
{
    return proxy_domain + prefix + constant[type][key]
}

const constant = {
    api: {
        resource:'api/'
    },
    project: {
        user_project: 'user-project'
    },
    module: {
        resource: 'module'
    },
    doc: {
        generate: 'doc/generate',
        resource : 'document'
    }
}
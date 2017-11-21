const proxy_domain = 'api/'
const prefix = 'api/'
export const getUri = (type,key ) =>
{
    return proxy_domain + prefix + constant[type][key]
}

const constant = {
    api: {
        api:'apis/'
    },
    project: {
        user_project: 'user-project'
    },
    doc: {
        generate: 'doc/generate'
    }
}
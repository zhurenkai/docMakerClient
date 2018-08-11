const proxy_domain = 'api/'
const prefix = 'api/'
const getUri = (type,key ) =>
{
    return proxy_domain + prefix + constant[type][key]
}
const constant = {
  client: {
    request: 'client-api/client/request'
  },
  api: {
    resource: 'api/'
  },
  project: {
    user_project: 'user-project',
    projects_with_api: 'projects-with-api'
  },
  module: {
    resource: 'module'
  },
  doc: {
    generate: 'doc/generate',
    resource: 'document',
    markdown: 'doc/markdown'
  },
  project_host: {
    resource: 'project-host/'
  },
  host:{
    resource: 'host/'
  },
  user: {
    info: 'user-info'
  },
  db_config: {
    resource: 'db-config'
  },
  db_comment: {
    import: 'import-db-comments',
    storeMany: 'comment/store-many',
    getHash: 'comment/hash'
  }
}


export { getUri }

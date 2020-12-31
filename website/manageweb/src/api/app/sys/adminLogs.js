import request from '@/utils/request'

export function adminlogs_request(token){
  return request({
    url: '/adminlogs/list',
    method: 'get',
    param: { token }
  })
}

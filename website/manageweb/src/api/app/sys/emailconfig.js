import request from '@/utils/request'

export function requestemailconfig(token){
  return request({
    url: '/emailconfig/list',
    method: 'get',
    param: { token }
  })
}

import request from '@/utils/request'

export function requestsysconfigs(token) {
  return request({
    url: '/systemconfig/list',
    method: 'get',
    params: { token }
  })
}

export function updatesysconfigs(data){
  return request({
    url: '/systemconfig/update',
    method: 'post',
    data
  })
}

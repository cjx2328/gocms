import request from '@/utils/request'

export function requestsysconfigs(token) {
  return request({
    url: '/systemconfig/list',
    method: 'get',
    params: { token }
  })
}

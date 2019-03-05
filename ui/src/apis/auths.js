import request from './request'

export const fetchAuth = (params = {}) => request.get('/oauth/servers', { params })
export const fetchAuthByName = (name, params = {}) => request.get(`/oauth/servers/${name}/`, { params })
export const deleteAuthByName = (name, params = {}) => request.delete(`/oauth/servers/${name}/`, { params })
export const putAuth = (name, payload) => request.put(`/oauth/servers/${name}/`, payload)

export const createAuth = (payload = {}) => request.post('/oauth/servers/', payload)

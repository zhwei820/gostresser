import request from './request'

export const StartStress = (id, payload) => request.post(`/start/${id}/`, payload)
export const StopStress = (id, payload) => request.post(`/stop/${id}/`, payload)


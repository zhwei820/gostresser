// const nodeEnv = process.env.VUE_APP_API
const MAIN_CONFIG = window.MAIN_CONFIG

// API Config
export const BACK_API =  MAIN_CONFIG === undefined ? 'http://localhost:8179': MAIN_CONFIG.gateway.GoStresser_api

// token
export const TOKEN_NAME = 'GOSTRESSER_TOKEN'
export const USER_INFO = 'GOSTRESSER_USER_INFO'


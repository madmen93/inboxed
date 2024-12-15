import type { IApiResponse, IEmail } from "@/types"
import { useDataStore } from "@/stores/dataStore"

const BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080'


export function useData() {
const dataStore = useDataStore()

    async function fetchData() {
        const response = await fetch(`${BASE_URL}/emails?page=${dataStore.page}&size=${dataStore.size}&sort=date&order=${dataStore.order}&field=${dataStore.field}&q=${dataStore.query}&fromDate=${dataStore.fromDate}&toDate=${dataStore.toDate} `, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        })
        if (response.ok) {
            const dataResponse = await response.json()
            const thisdata = dataResponse as IApiResponse
            dataStore.setData(thisdata)
          }else{
            console.error('Error al obtener datos:', response.statusText)
            return
          }
    
    }

    async function getEmailByID() {
        const response = await fetch(`${BASE_URL}/emails/${dataStore.id}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
        })
        
        if (response.ok) {
            const dataResponse = await response.json()
            const thisdata = dataResponse as IEmail
            dataStore.setCurrentEmail(thisdata)
        }else{
            console.error('Error al obtener datos:', response.statusText)
            return
        }
    }
    return { fetchData, getEmailByID, dataStore }
}
import { defineStore } from "pinia";
import { ref, computed } from "vue";
import type { IApiResponse,  IEmail } from "@/types";

export const useDataStore = defineStore('data', () => {
    const data = ref<IApiResponse>()
    const currentEmail = ref<IEmail|null>(null)
    const page = ref(1)
    const size = ref(10)
    const query = ref('')
    const messageID = ref('')
    const showEmail = ref(false)
    const id = ref('')
    const field = ref('content')
    const order = ref('desc')
    const showFilters = ref(false)
    const fromDate = ref('')
    const toDate = ref('')

    const totalEmails = computed(() => data.value?.hits.total.value || 0);
    const totalPages = computed(() => Math.ceil(totalEmails.value /size.value))

    function toggleEmail() {
        showEmail.value = !showEmail.value
    }

    function changePage(increment: boolean) {
        if (increment && totalPages.value > page.value) {
            page.value++
        }else if(!increment && page.value > 1) {
            page.value--
        }

        window.scrollTo({
            top: 0,
            behavior: 'smooth' 
        });
    }

    function setMessageID(newMessageID: string) {
        messageID.value = newMessageID
        showEmail.value = true
    }

    function setID(newID: string) {
        id.value = newID
        showEmail.value = true
    }

    function setData(newData: IApiResponse) {
        data.value = newData
    }

    function setCurrentEmail(newEmail: IEmail) {
        currentEmail.value = newEmail; 
    }

    function setFiltersView(value: boolean) {
        showFilters.value = value
    }

    function resetFilters() {
        order.value = 'desc'
        field.value = 'content'
        query.value = ''
        size.value = 10
        page.value = 1
    }

    return { query, totalPages, changePage, data, setData, totalEmails, page, size, setMessageID, messageID, setCurrentEmail, currentEmail, toggleEmail, showEmail, id, setID, field, order, showFilters, setFiltersView, resetFilters, fromDate, toDate}
})
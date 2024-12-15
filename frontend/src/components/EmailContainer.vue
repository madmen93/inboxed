<script lang="ts" setup>
import { onMounted, watch } from 'vue';
import { useData } from '@/services/email';
import { useDataStore } from '@/stores/dataStore';

const dataStore = useDataStore()
const { fetchData, getEmailByID} = useData()

async function showEmail(messageID: string) {
    dataStore.setID(messageID);
    getEmailByID();
}

watch(() => [dataStore.page, dataStore.size, dataStore.query], () => {
  fetchData();
});

onMounted(async () => {
  try {
    await fetchData()
  } catch (error) {
    console.error('Error al cargar los datos:', error)
  }
})
</script>

<template>
    <section class="flex flex-col bg-amber-50">
        <form class="flex justify-center gap-2" @submit.prevent>
            <input class="p-4 rounded-md text-emerald-900 placeholder:italic sm:w-2/5 md:w-1/4 shadow-sm border border-slate-300 focus:border-emerald-500 focus:ring-emerald-500 focus:ring-1" v-model="dataStore.query" type="text" placeholder="Start searching here...">
            <select class="p-4 rounded-md bg-emerald-600 text-white shadow-sm" v-model="dataStore.size">
                <option value="10">10</option>
                <option value="25">25</option>
                <option value="50">50</option>
                <option value="75">75</option>
                <option value="100">100</option>
            </select>
        </form>

        <div  class="self-center w-4/5 sm:w-[630px] md:w-3/4 lg:w-2/4 pt-5 flex justify-between items-center" v-if="dataStore.data && dataStore.data.hits && dataStore.data.hits.hits.length !== 0">
            <p class="font-bold italic text-emerald-600">There are {{ dataStore.data?.hits.total.value }} results...</p>
            <button class="bg-emerald-600 text-white font-bold p-3 rounded-xl hover:bg-emerald-500 focus:outline-none focus:ring focus:ring-emerald-800" @click="dataStore.setFiltersView(true)">Advanced Filters</button>
        </div>
        <!-- Contenido de emails -->
        <section class="flex justify-center w-full">
            <div class=" bg-white mt-5 mb-6 rounded-md w-4/5 sm:w-[630px] md:w-3/4 lg:w-2/4 p-4 text-slate-400" v-if="dataStore.data && dataStore.data.hits && dataStore.data.hits.hits.length === 0">
                <p>Sorry, no results found...</p>
            </div>
            <div class=" bg-white mt-5 mb-6 rounded-md w-4/5 sm:w-[630px] md:w-3/4 lg:w-2/4"  v-if="dataStore.data && dataStore.data.hits && dataStore.data.hits.hits.length > 0">
                <ul class="p-3 rounded-md flex flex-col gap-8 " >
                    <li class="flex flex-col gap-2 sm:gap-0 sm:flex-row sm:justify-between sm:items-center" v-for="(email, index) in dataStore.data.hits.hits" :key="index">
                        <div class="w-full ">
                            <div class="flex flex-col sm:flex-row sm:gap-3">
                                <div class="min-w-14 md:min-w-24 text-left">
                                    <label class="font-bold text-emerald-600">Subject</label>
                                </div>
                                <div class="flex-grow break-words ">
                                    <p class="pl-2 sm:pl-0">{{ email._source.subject }}</p>
                                </div>
                            </div>
                            <div class="flex flex-col sm:flex-row sm:gap-3">
                                <div class="min-w-14 md:min-w-24 text-left">
                                    <label class="font-bold text-emerald-600">From</label>
                                </div>
                                <div class="flex-grow break-words sm:truncate min-w-0">
                                    <p class="pl-2 sm:pl-0">{{ email._source.from }}</p>
                                </div>
                            </div>
                            <div class="flex flex-col sm:flex-row sm:gap-3">
                                <div class="min-w-14 md:min-w-24 text-left">
                                    <label class="font-bold text-emerald-600">Date</label>
                                </div>
                                    <div class="flex-grow break-words sm:truncate min-w-0">
                                    <p class="pl-2 sm:pl-0">{{ email._source.date }}</p>
                                </div>
                            </div>
                        </div>
                        <div class="flex justify-center mb-3 sm:mb-0">
                            <button class="bg-emerald-600 text-white font-bold p-3 rounded-xl hover:bg-emerald-500 focus:outline-none focus:ring focus:ring-emerald-800" @click="() => showEmail(email._id)">More</button>
                        </div>
                        <hr>
                    </li>
                </ul>
            </div>
        </section>
        <!-- Botones de páginas -->
        <div class="flex justify-center items-center gap-4 pb-5" v-if="dataStore.data && dataStore.data.hits && dataStore.data.hits.hits.length !== 0">
            <button :class="[
            'rounded-xl',
            'p-2',
            'w-20',
            'font-bold',
            dataStore.page === 1 ? 'bg-emerald-800' : 'bg-emerald-600',
            dataStore.page === 1 ? 'text-gray-300' : 'text-white',
            dataStore.page === 1 ? 'cursor-not-allowed' : 'hover:bg-emerald-500'
          ]" @click="()=>dataStore.changePage(false)" :disabled="dataStore.page === 1">Previous</button>
            <span>Page {{ dataStore.page }} of {{ dataStore.totalPages }}</span>
            <button :class="[
            'rounded-xl',
            'p-2',
            'w-20',
            'font-bold',
            dataStore.page === dataStore.totalPages || dataStore.page > dataStore.totalPages ? 'bg-emerald-800' : 'bg-emerald-600',
            dataStore.page === dataStore.totalPages || dataStore.page > dataStore.totalPages ? 'text-gray-300' : 'text-white',
            dataStore.page === dataStore.totalPages || dataStore.page > dataStore.totalPages ? 'cursor-not-allowed' : 'hover:bg-emerald-500'
          ]" @click="()=>dataStore.changePage(true)" :disabled="dataStore.page === dataStore.totalPages || dataStore.page > dataStore.totalPages">Next</button>
        </div>
    </section>
</template>
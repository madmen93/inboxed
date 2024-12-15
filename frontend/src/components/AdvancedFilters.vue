<script setup lang="ts">
    import {ref} from 'vue'
    import { useDataStore } from '@/stores/dataStore';
    import { useData } from '@/services/email';

    const searchByRange = ref<boolean>()
    const {fetchData} = useData()

    const dataStore = useDataStore()

    function applyFilter() {
        fetchData()
        dataStore.setFiltersView(false)
    }

</script>

<template>
    <section class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50 transition-opacity duration-300">
        <section class="bg-white flex flex-col justify-start items-center rounded-md shadow-lg relative w-[90vw] sm:w-[70vw] md:w-[40vw] max-w-3xl h-[70vh] md:h-[43vh] overflow-hidden">
            <div class="bg-emerald-600 py-3 px-4 w-full mb-6">
                <h2 class="text-white font-bold text-center">Advanced Filters</h2>
            </div>
            <form class="flex flex-col gap-4 w-4/5" @submit.prevent="applyFilter">
                <div class="flex justify-between  w-full">
                    <label class="text-emerald-600 font-bold text-lg">Sorted by date:</label>
                    <select class="p-2 rounded-md bg-gray-300 text-black shadow-sm" v-model="dataStore.order">
                        <option value="desc">Desc</option>
                        <option value="asc">Asc</option>
                    </select>
                </div>
                <div class="flex justify-between w-full">
                    <label class="text-emerald-600 font-bold text-lg">Search by field:</label>
                    <select class="p-2 rounded-md bg-gray-300 text-black shadow-sm" v-model="dataStore.field">
                        <option value="content">Content</option>
                        <option value="subject">Subject</option>
                        <option value="from">From</option>
                        <option value="to">To</option>
                    </select>
                </div>
                <div class="bg-amber-50">
                    <p class="p-4"><span class="font-bold">Note:</span> To search by 'from' or 'to', you need to omit the part of the email after the '@'. For example, to search for 'name_xy@mail.com', use 'name_xy'.</p>
                </div>
                <div class="flex justify-between w-full">
                    <button @click="dataStore.setFiltersView(false)" name="reset" class="bg-red-600 text-white font-bold p-3 rounded-xl hover:bg-red-500 focus:outline-none focus:ring focus:ring-red-800">Cancel</button>
                    <button type="submit" name="apply" class="bg-emerald-600 text-white font-bold p-3 rounded-xl hover:bg-emerald-500 focus:outline-none focus:ring focus:ring-emerald-800">Apply</button>
                </div>
            </form>
        </section>
    </section> 
</template>
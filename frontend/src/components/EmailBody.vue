<script lang="ts" setup>
    import { useDataStore } from '@/stores/dataStore';

    const dataStore = useDataStore()

    const splitContent = (text: string): string[] => { 
        const query = dataStore.query || '';
        if (!query) return [text]; 

        const regex = new RegExp(`(${query})`,'gi')

        return text.split(regex)
    }
    
    function isQuery(chunk:string): boolean {
        return chunk.toLowerCase() === (dataStore.query || '').toLowerCase();
    }

</script>

<template>
    <section class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50 transition-opacity duration-300">
        <section class="bg-white flex flex-col justify-center items-center rounded-md shadow-lg relative w-[90vw] max-w-3xl h-[80vh] overflow-hidden">
            <div class="bg-emerald-600 py-3 px-4 w-full">
                <p class="text-white text-center text-xs"><span class="font-bold text-sm">Message_ID:</span> {{ dataStore.currentEmail?._source.message_id }}</p>
            </div>
            <section class="w-full px-6 py-4 overflow-y-auto flex-1" >
                <div>
                    <label class="text-emerald-600 font-bold text-lg">Date:</label>
                    <p> {{ dataStore.currentEmail?._source.date }}</p>
                </div>
                <div>
                    <label class="text-emerald-600 font-bold text-lg">From:</label>
                    <p> {{ dataStore.currentEmail?._source.from }}</p>
                </div>
                <div>
                    <label class="text-emerald-600 font-bold text-lg">To:</label>
                    <p> {{ dataStore.currentEmail?._source.to }}</p>
                </div>
                <div>
                    <label class="text-emerald-600 font-bold text-lg">Subject:</label>
                    <p> {{ dataStore.currentEmail?._source.subject }}</p>
                </div>
                <div>
                    <label class="text-emerald-600 font-bold text-lg">X_from:</label>
                    <p> {{ dataStore.currentEmail?._source.x_from }}</p>
                </div>
                <div>
                    <label class="text-emerald-600 font-bold text-lg">X_to:</label>
                    <p> {{ dataStore.currentEmail?._source.x_to }}</p>
                </div>
                <div>
                    <label class="text-emerald-600 font-bold text-lg">X_bcc:</label>
                    <p> {{ dataStore.currentEmail?._source.x_bcc }}</p>
                </div>
                <div>
                    <label class="text-emerald-600 font-bold text-lg">X_cc:</label>
                    <p> {{ dataStore.currentEmail?._source.x_cc }}</p>
                </div>
                <div>
                    <label class="text-emerald-600 font-bold text-lg">X_filename:</label>
                    <p> {{ dataStore.currentEmail?._source.x_filename }}</p>
                </div>
                <div>
                    <label class="text-emerald-600 font-bold text-lg">X_folder:</label>
                    <p> {{ dataStore.currentEmail?._source.x_folder }}</p>
                </div>
                <div>
                    <label class="text-emerald-600 font-bold text-lg">Content:</label>
                    <p v-if="dataStore.query !=='' && dataStore.currentEmail?._source.content !== undefined && dataStore.currentEmail?._source.content !== null">
                        <span v-for="(chunk, index) in splitContent(dataStore.currentEmail._source.content)" :key="index" :class="{ 'bg-emerald-300 text-black': isQuery(chunk) }">
                            {{ chunk }}
                        </span>
                    </p>
                    <p v-else>{{ dataStore.currentEmail?._source.content }}</p>
                </div>
            </section>
            <div class="w-full py-3 flex justify-center border-t border-gray-200 px-6">
                <button class="bg-emerald-600 text-white font-bold p-2 w-20 rounded-xl hover:bg-emerald-500" @click="dataStore.toggleEmail">
                    Close
                </button>
            </div>
        </section>
    </section>
</template>
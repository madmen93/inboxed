<script setup lang="ts">
import EmailContainer from '@/components/EmailContainer.vue';
import EmailBody from '@/components/EmailBody.vue';
import AdvancedFilters from '@/components/AdvancedFilters.vue';
import { useDataStore } from '@/stores/dataStore';

const dataStore = useDataStore()

</script>

<template>
  <main class="w-full h-screen bg-amber-50">
    <header class="bg-emerald-600">
      <div class="h-32 flex justify-center items-end pb-4">
        <h1 class="text-amber-50 text-5xl">Inboxed</h1>
      </div>
    </header>
    <EmailContainer class="mt-8 mb-6"/>
    <EmailBody v-if="dataStore.showEmail"/>
    <AdvancedFilters v-if="dataStore.showFilters"/>
  </main>
</template>

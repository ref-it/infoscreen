<template>
    <div class="bg-black/20 text-xl font-bold p-4 border-b border-zinc-700 pl-4 pr-2 py-2 rounded-t-lg grid header-with-badge">
        <span>Veranstaltungen</span>
    </div>
    <div v-if="events.length > 0" class="h-full grid grid-box3 overflow-y-auto">
        <div v-if="events.length > 0">
            <div class="h-full w-full">
                <div v-for="(item, index) in events" :key="index" class="grid grid-events border-b border-zinc-700">
                    <div class="px-4 py-2">
                        <div class="text-base">{{ weekday[new Date(item.startDate).getDay()] }}, {{ item.startDate.split('-')[2] }}.{{ item.startDate.split('-')[1] }}.</div>
                        <div class="text-base">{{ item.startTime }}</div>
                    </div>
                    <div class="px-4 py-2 text-base">{{ item.summary }}</div>
                </div>
            </div>
        </div>
        <div v-else class="px-4 py-2 text-xl m-auto text-center">
            <p>Keine Veranstaltungen in nächster Zeit.</p>
        </div>
    </div>
</template>
  
<script setup lang="ts">
    import { ref } from 'vue'
    import axios from 'axios'
    
    // Reactive list of departures used in the view
    const events: any = ref([])

    const weekday = ["So","Mo","Di","Mi","Do","Fr","Sa"];

    
    const getEvents = () => {
        const url = '/api/events'
        axios
            .get(url)
            .then((response: any) => {
                events.value = response.data
            })
    }
    
    // request data periodically
    getEvents()
    setInterval(() => {
        getEvents()
    }, 60000)
</script>
  
<style scoped>
    .grid-box3 {
        grid-template-rows: auto 1fr;
    }
    .grid-events {
        grid-template-columns: 9rem 1fr;
    }
</style>
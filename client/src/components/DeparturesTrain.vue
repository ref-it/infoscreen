<template>
    <div class="bg-black/20 text-xl font-bold p-4 border-b border-zinc-700 pl-4 pr-2 py-2 rounded-t-lg grid header-with-badge">
        <span>Abfahrten</span>
        <div class="-mt-02">
            <span class="inline-flex items-center rounded-md bg-emerald-400/10 ml-auto px-2 py-1 text-sm text-emerald-400 ring-1 ring-inset ring-emerald-400/30">{{ stationName }}</span>
        </div>
    </div>
    <div v-if="trainDepartures.length > 0" class="h-full grid grid-box3">
        <div v-if="trainDepartures.length > 0">
            <div class="grid grid-train bg-black/20 border-b border-zinc-700">
                <div class="px-4 py-1.5 text-sm font-bold">Linie</div>
                <div class="px-4 py-1.5 text-sm font-bold">Abfahrt</div>
                <div class="px-4 py-1.5 text-sm font-bold">Ziel</div>
                <div class="px-4 py-1.5 text-sm font-bold">Gleis</div>
            </div>
            <div class="h-full w-full overflow-y-auto">
                <div v-for="(connection, index) in trainDepartures" :key="index" class="grid grid-train border-b border-zinc-700">
                    <div class="px-4 py-2 text-base">{{ connection.category }} {{ connection.line }}</div>
                    <div class="px-4 py-2 text-base">{{ connection.hours }}:{{ connection.minutes }}</div>
                    <div class="px-4 py-2 text-base">{{ connection.destination }}</div>
                    <div class="px-4 py-2 text-base">{{ connection.track }}</div>
                </div>
            </div>
        </div>
        <div v-else class="px-4 py-2 text-xl m-auto text-center p-4">
            <p>Keine Verbindungen in nächster Zeit.</p>
        </div>
    </div>
</template>
  
<script setup lang="ts">
    import { ref } from 'vue'
    import axios from 'axios'

    const props = defineProps({
        id: String
    })
    
    // Reactive list of departures used in the view
    const trainDepartures: any = ref([])

    // Reactive station name
    const stationName = ref('')
    
    const getTrainDepartures = () => {
        const url = '/api/departures/train?station=' + props.id
        axios
            .get(url)
            .then((response: any) => {
                stationName.value = response.data.stationName
                trainDepartures.value = response.data.departures
            })
    }
    
    // request data periodically
    getTrainDepartures()
    setInterval(() => {
        getTrainDepartures()
    }, 60000)
</script>
  
<style scoped>
    .grid-box3 {
        grid-template-rows: auto 1fr;
    }
    .grid-train {
        grid-template-columns: 6.8rem 5.3rem 1fr auto;
    }
</style>
<template>
    <div class="bg-black/20 text-xl font-bold border-b border-slate-700 px-4 py-2 rounded-t-lg">Busverbindungen (G.-Kirchhoff-Platz)</div>
    <div class="h-full grid-box2 pb-2 flex">
        <div v-if="weekday != 0 && weekday != 6 && connections_bus_shown.length > 0" class="w-full">
            <div class="grid grid-bus bg-black/20 border-b border-slate-700">
                <div class="px-4 py-1.5 text-sm font-bold">Linie</div>
                <div class="px-4 py-1.5 text-sm font-bold">Abfahrt G.-Kirchhoff-Platz</div>
                <div class="px-4 py-1.5 text-sm font-bold">Ankunft Bahnhof</div>
            </div>
            <div v-for="(connection, index) in connections_bus_shown" :key="index" class="grid grid-bus border-b border-slate-700">
                <div class="px-4 py-2 text-xl">{{ connection.line }}</div>
                <div class="px-4 py-2 text-xl">{{ connection.start }}</div>
                <div class="px-4 py-2 text-xl">{{ connection.end }}</div>
            </div>
        </div>
        <div v-else-if="weekday != 0 && weekday != 6 && connections_bus_shown == 0" class="px-4 py-2 text-xl m-auto text-center">
            <p>Keine Verbindungen in nächster Zeit.</p>
        </div>
        <div v-else class="px-3 py-2 text-xl m-auto m-auto text-center">
            <p>CAMIL (Linie D) fährt nur von Montag bis Freitag.</p>
        </div>
        <div class="px-3 mt-auto">
            <!-- <div class="w-full h-px bg-slate-700 mb-2" aria-hidden="true"></div> -->
            <img class="w-full rounded-lg max-h-52" src="@/assets/img/camil-haltestelle-kirchhoff.svg">
        </div>
    </div>
</template>
  
<script setup lang="ts">
    import { ref } from 'vue'

    const weekday = ref()

    const getCurrentDateTime = () => {
        const date = new Date();
        weekday.value = date.getDay();
    }
  
    const connections_bus = [
        {
            line: "D",
            start: "10:05",
            end: "10:20"
        },
        {
            line: "D",
            start: "11:05",
            end: "11:20"
        },
        {
            line: "D",
            start: "12:05",
            end: "12:20"
        },
        {
            line: "D",
            start: "13:55",
            end: "14:10"
        },
        {
            line: "D",
            start: "14:55",
            end: "15:10"
        },
        {
            line: "D",
            start: "15:55",
            end: "16:10"
        },
        {
            line: "D",
            start: "16:55",
            end: "17:10"
        }
    ]
  
    const connections_bus_shown = ref();
    
    const fillConnectionsBus = () => {
        connections_bus_shown.value = []
        const date = new Date()
        const hour = date.getHours()
        const minute = date.getMinutes()
    
        for (let i = 0; i < connections_bus.length; i++) {
            if (connections_bus_shown.value.length < 2) {
                const connectionHour = +connections_bus[i].start.split(':')[0]
                const connectionMinute = +connections_bus[i].start.split(':')[1]
                if (hour <= connectionHour && connectionHour - hour <= 2) {
                    if (minute <= connectionMinute) {
                        connections_bus_shown.value.push(connections_bus[i])
                    }
                }
            }
        }
    }
  
    getCurrentDateTime()
    fillConnectionsBus()
    setInterval(() => {
        getCurrentDateTime()
        fillConnectionsBus()
    }, 1000)
</script>
  
<style scoped>
    .grid-box2 {
        display: grid;
        grid-template-rows: 1fr auto;
    }

    .grid-bus {
        grid-template-columns: 1fr 13.5rem 10rem;
    }
</style>
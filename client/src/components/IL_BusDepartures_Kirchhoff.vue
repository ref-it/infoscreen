<template>

        <div class="bg-black/20 text-xl font-bold p-4 border-b border-zinc-700 pl-4 pr-2 py-2 rounded-t-lg grid header-with-badge">
        <span>Abfahrten</span>
        <div class="-mt-02">
            <span class="inline-flex items-center rounded-md bg-emerald-400/10 ml-auto px-2 py-1 text-sm text-emerald-400 ring-1 ring-inset ring-emerald-400/30">G.-Kirchhoff-Platz</span>
        </div>
    </div>
    <div class="h-full grid grid-box3">
        <div v-if="weekday != 0 && weekday != 6 && connections_bus_shown.length > 0">
            <div class="grid grid-train bg-black/20 border-b border-zinc-700">
                <div class="px-4 py-1.5 text-sm font-bold">Linie</div>
                <div class="px-4 py-1.5 text-sm font-bold">Abfahrt</div>
                <div class="px-4 py-1.5 text-sm font-bold">Ziel</div>
            </div>
            <div class="h-full w-full overflow-y-auto">
                <div v-for="(connection, index) in connections_bus_shown" :key="index" class="grid grid-train border-b border-zinc-700">
                    <div class="px-4 py-2 text-base">{{ connection.category }} {{ connection.line }}</div>
                    <div class="px-4 py-2 text-base">{{ connection.hours }}:{{ connection.minutes }}</div>
                    <div class="px-4 py-2 text-base">{{ connection.destination }}</div>
                </div>
            </div>
        </div>
        <div v-else-if="weekday != 0 && weekday != 6 && connections_bus_shown == 0" class="px-4 py-2 text-xl m-auto text-center">
            <p>Keine Verbindungen in nächster Zeit.</p>
        </div>
        <div v-else class="px-3 py-2 text-xl m-auto m-auto text-center">
            <p>CAMIL (Linie D) fährt nur von Montag bis Freitag.</p>
        </div>
        <!--<div class="px-3 mt-auto">-->
            <!-- <div class="w-full h-px bg-zinc-700 mb-2" aria-hidden="true"></div> -->
            <!--<img class="w-full rounded-lg max-h-52" src="@/assets/img/camil-haltestelle-kirchhoff.svg">
        </div>-->
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
            category: "Bus",
            line: "D",
            hours: "10",
            minutes: "05",
            destination: "Neuhäuser Weg"
        },
        {
            category: "Bus",
            line: "D",
            hours: "11",
            minutes: "05",
            destination: "Neuhäuser Weg"
        },
        {
            category: "Bus",
            line: "D",
            hours: "12",
            minutes: "05",
            destination: "Neuhäuser Weg"
        },
        {
            category: "Bus",
            line: "D",
            hours: "13",
            minutes: "55",
            destination: "Neuhäuser Weg"
        },
        {
            category: "Bus",
            line: "D",
            hours: "14",
            minutes: "55",
            destination: "Neuhäuser Weg"
        },
        {
            category: "Bus",
            line: "D",
            hours: "15",
            minutes: "55",
            destination: "Neuhäuser Weg"
        },
        {
            category: "Bus",
            line: "D",
            hours: "16",
            minutes: "55",
            destination: "Neuhäuser Weg"
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
                if (hour <= +connections_bus[i].hours && +connections_bus[i].hours - hour <= 2) {
                    if (minute <= +connections_bus[i].minutes) {
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
    .grid-box3 {
        grid-template-rows: auto 1fr;
    }
    .grid-train {
        grid-template-columns: 6.8rem 5.3rem 1fr;
    }
</style>
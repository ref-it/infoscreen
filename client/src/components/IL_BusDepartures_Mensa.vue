<template>
    <div class="bg-black/20 text-xl font-bold p-4 border-b border-zinc-700 pl-4 pr-2 py-2 rounded-t-lg grid header-with-badge">
        <span>Busverbindungen</span>
        <div class="-mt-02">
            <span class="inline-flex items-center rounded-md bg-blue-400/10 ml-auto px-2 py-1 text-sm text-blue-400 ring-1 ring-inset ring-blue-400/30">Mensa</span>
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
                    <div class="px-4 py-2 text-xl">{{ connection.category }} {{ connection.line }}</div>
                    <div class="px-4 py-2 text-xl">{{ connection.hours }}:{{ connection.minutes }}</div>
                    <div class="px-4 py-2 text-xl">{{ connection.destination }}</div>
                </div>
            </div>
        </div>
        <div v-else class="w-full h-full flex items-center justify-center text-xl p-4">
            <p>Keine Verbindungen in nächster Zeit.</p>
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
            category: "Bus",
            line: "A",
            hours: "04",
            minutes: "45",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "A",
            hours: "05",
            minutes: "15",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "A",
            hours: "05",
            minutes: "45",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "A",
            hours: "06",
            minutes: "15",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "A",
            hours: "06",
            minutes: "45",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "A",
            hours: "07",
            minutes: "15",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "A",
            hours: "07",
            minutes: "45",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "A",
            hours: "08",
            minutes: "15",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "A",
            hours: "08",
            minutes: "45",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "A",
            hours: "09",
            minutes: "15",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "A",
            hours: "09",
            minutes: "45",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "D",
            hours: "09",
            minutes: "52",
            destination: "G.-Kirchhoff-Platz"
        },
        {
            category: "Bus",
            line: "D",
            hours: "10",
            minutes: "13",
            destination: "Neuhäuser Weg"
        },
        {
            category: "Bus",
            line: "A",
            hours: "10",
            minutes: "15",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "A",
            hours: "10",
            minutes: "45",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "D",
            hours: "10",
            minutes: "52",
            destination: "G.-Kirchhoff-Platz"
        },
        {
            category: "Bus",
            line: "D",
            hours: "11",
            minutes: "13",
            destination: "Neuhäuser Weg"
        },
        {
            category: "Bus",
            line: "A",
            hours: "11",
            minutes: "15",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "A",
            hours: "11",
            minutes: "45",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "D",
            hours: "11",
            minutes: "52",
            destination: "G.-Kirchhoff-Platz"
        },
        {
            category: "Bus",
            line: "D",
            hours: "12",
            minutes: "13",
            destination: "Neuhäuser Weg"
        },
        {
            category: "Bus",
            line: "A",
            hours: "12",
            minutes: "15",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "A",
            hours: "12",
            minutes: "45",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "A",
            hours: "13",
            minutes: "15",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "D",
            hours: "13",
            minutes: "42",
            destination: "G.-Kirchhoff-Platz"
        },
        {
            category: "Bus",
            line: "A",
            hours: "13",
            minutes: "45",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "D",
            hours: "14",
            minutes: "03",
            destination: "Neuhäuser Weg"
        },
        {
            category: "Bus",
            line: "A",
            hours: "14",
            minutes: "15",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "D",
            hours: "14",
            minutes: "42",
            destination: "G.-Kirchhoff-Platz"
        },
        {
            category: "Bus",
            line: "A",
            hours: "14",
            minutes: "45",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "D",
            hours: "15",
            minutes: "03",
            destination: "Neuhäuser Weg"
        },
        {
            category: "Bus",
            line: "A",
            hours: "15",
            minutes: "15",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "D",
            hours: "15",
            minutes: "42",
            destination: "G.-Kirchhoff-Platz"
        },
        {
            category: "Bus",
            line: "A",
            hours: "15",
            minutes: "45",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "D",
            hours: "16",
            minutes: "03",
            destination: "Neuhäuser Weg"
        },
        {
            category: "Bus",
            line: "A",
            hours: "16",
            minutes: "15",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "D",
            hours: "16",
            minutes: "42",
            destination: "G.-Kirchhoff-Platz"
        },
        {
            category: "Bus",
            line: "A",
            hours: "16",
            minutes: "45",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "D",
            hours: "17",
            minutes: "03",
            destination: "Neuhäuser Weg"
        },
        {
            category: "Bus",
            line: "A",
            hours: "17",
            minutes: "15",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "A",
            hours: "17",
            minutes: "45",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "A",
            hours: "18",
            minutes: "15",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "A",
            hours: "18",
            minutes: "45",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "A",
            hours: "19",
            minutes: "15",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "A",
            hours: "19",
            minutes: "45",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "A",
            hours: "20",
            minutes: "15",
            destination: "Eichicht"
        },
        {
            category: "Bus",
            line: "A",
            hours: "21",
            minutes: "15",
            destination: "Eichicht"
        },
    ]
  
    const connections_bus_shown = ref();
    
    const fillConnectionsBus = () => {
        connections_bus_shown.value = []
        const date = new Date()
        const hour = date.getHours()
        const minute = date.getMinutes()
    
        for (let i = 0; i < connections_bus.length; i++) {
                const connectionHour = +connections_bus[i].hours
                const connectionMinute = +connections_bus[i].minutes
                if (hour < connectionHour && connectionHour - hour <= 3 || hour == connectionHour && minute <= connectionMinute) {
                    connections_bus_shown.value.push(connections_bus[i])
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
        grid-template-columns: 6.4rem 5.5rem 1fr;
    }
</style>

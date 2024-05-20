<template>
    <div class="bg-black/20 text-xl font-bold p-4 border-b border-zinc-700 px-4 py-2 rounded-t-lg">Hat das Café offen?</div>
    <div class="flex flex-1 self-stretch w-full h-full overflow-auto">
        <div v-if="open" class="w-full h-full flex flex-col items-center justify-center">
            <p class="text-xl font-bold">{{ eventName }}</p>
            <p class="text-base mt-2">{{ eventStartTime }} &ndash; {{ eventEndTime }} Uhr</p>
            <p v-if="eventMarquee != undefined" class="text-base mt-6">{{ eventMarquee }}</p>
        </div>
        <div v-else class="w-full h-full flex items-center justify-center text-xl text-center">
            Heute hat das Café leider geschlossen.
        </div>
    </div>
</template>
<script setup lang="ts">
    import { ref } from 'vue'
    import axios from 'axios'

    const open = ref(false)
    const eventName = ref()
    const eventMarquee = ref()
    const eventStartTime = ref()
    const eventEndTime = ref()

    const getCafeOpenings = () => {
        const url = '/api/bc-cafe'
        axios
            .get(url)
            .then((response: any) => {
                const date = new Date()
                const currentDateStringISO =  date.getFullYear() + '-' + ("0" + (date.getMonth() + 1)).slice(-2) + '-' + ("0" + date.getDate()).slice(-2)
                for (let i = 0; i < response.data.events.length; i++) {
                    if (response.data.events[i].start == currentDateStringISO && response.data.events[i].cancelled == false) {
                        open.value = true
                        eventName.value = response.data.events[i].name
                        eventMarquee.value = response.data.events[i].marquee
                        eventStartTime.value = response.data.events[i].start_time.substring(0, response.data.events[i].start_time.length-3)
                        eventEndTime.value = response.data.events[i].end_time.substring(0, response.data.events[i].end_time.length-3)
                    }
                }
            })
            .catch(() => {
                open.value = false
            })
    }

    getCafeOpenings()
    setInterval(() => {
        getCafeOpenings()
    }, 300000)
</script>

<template>
    <div class="bg-black/20 text-xl font-bold p-4 border-b border-zinc-700 px-4 py-2 rounded-t-lg">Warnungen</div>
    <div class="flex flex-1 self-stretch w-full h-full overflow-auto">
        <div v-if="warnings[0] != undefined" class="w-full text-base">
            <div v-for="(item, index) in warnings" :key="index" class="px-4 py-2 border-b border-zinc-700">
                <h2 class="pb-1 font-bold">{{ item.info[0].headline }}</h2>
                <p>{{ item.info[0].description }}</p>
            </div>
        </div>
        <div v-else class="w-full h-full flex items-center justify-center text-xl">
            Es liegen keine Warnungen vor.
        </div>
    </div>
</template>

<script setup lang="ts">
    import { ref } from 'vue'
    import axios from 'axios'

    const props = defineProps({
        id: String
    })

    const warnings = ref()

    const getWeatherForecast = () => {
        const url = '/api/warnings?location=' + props.id
        axios
            .get(url)
            .then((response: any) => {
                console.log(response.data)
                warnings.value = response.data
            })
            .catch(() => {
                warnings.value = []
            })
    }

    getWeatherForecast()
    setInterval(() => {
        getWeatherForecast()
    }, 300000)
</script>

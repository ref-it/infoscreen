<template>
    <div class="bg-black/20 text-xl font-bold p-4 border-b border-slate-700 px-4 py-2 rounded-t-lg">Wetter</div>
    <div class="flex flex-1 self-stretch w-full h-full overflow-auto">
        <div class="w-full grid grid-cols-4 grid-rows-[3rem_1fr_3rem_3rem] gap-x-4 gap-y-0.5 text-center p-4 text-base">
            <div v-for="(item, index) in forecastDates" :key="index" class="bg-black/20 rounded-t-md flex items-center justify-center font-bold">{{ item }}</div>

            <div v-for="(item, index) in forecast.days" :key="index" class="bg-black/20 p-2 flex items-center justify-center">
                <img v-if="item.icon1 == 1" src="/weather-icons/svg/wi-day-sunny.svg">
                <img v-if="item.icon1 == 2" src="/weather-icons/svg/wi-day-sunny-overcast.svg">
                <img v-if="item.icon1 == 3" src="/weather-icons/svg/wi-day-cloudy.svg">
                <img v-if="item.icon1 == 4" src="/weather-icons/svg/wi-cloudy.svg">
                <img v-if="item.icon1 == 5 || item.icon1 == 6" src="/weather-icons/svg/wi-dust.svg">
                <img v-if="item.icon1 == 7 || item.icon1 == 8 || item.icon1 == 10 || item.icon1 == 11" src="/weather-icons/svg/wi-showers.svg">
                <img v-if="item.icon1 == 9" src="/weather-icons/svg/wi-rain.svg">
                <img v-if="item.icon1 == 12 || item.icon1 == 13" src="/weather-icons/svg/wi-rain-mix.svg">
                <img v-if="item.icon1 == 14 || item.icon1 == 15 || item.icon1 == 16" src="/weather-icons/svg/wi-snow-wind.svg">
                <img v-if="item.icon1 == 17" src="/weather-icons/svg/wi-hail.svg">
                <img v-if="item.icon1 == 18" src="/weather-icons/svg/wi-day-showers.svg">
                <img v-if="item.icon1 == 19" src="/weather-icons/svg/wi-day-rain.svg">
                <img v-if="item.icon1 == 20 || item.icon1 == 21" src="/weather-icons/svg/wi-day-rain-mix.svg">
                <img v-if="item.icon1 == 22 || item.icon1 == 23" src="/weather-icons/svg/wi-day-snow.svg">
                <img v-if="item.icon1 == 24 || item.icon1 == 25" src="/weather-icons/svg/wi-day-hail.svg">
                <img v-if="item.icon1 == 26 || item.icon1 == 27 || item.icon1 == 29" src="/weather-icons/svg/wi-storm-showers.svg">
                <img v-if="item.icon1 == 28 || item.icon1 == 30" src="/weather-icons/svg/wi-thunderstorm.svg">
                <img v-if="item.icon1 == 31" src="/weather-icons/svg/wi-strong-wind.svg">
            </div>

            <div v-for="(item, index) in forecast.days" :key="index" class="flex w-full">
                <span v-if="item.temperatureMax / 10 < -5" class="bg-sky-700 flex items-center justify-center w-full">{{ item.temperatureMax / 10 }} °C</span>
                <span v-if="item.temperatureMax / 10 < 0" class="bg-cyan-500 text-slate-800 flex items-center justify-center w-full">{{ item.temperatureMax / 10 }} °C</span>
                <span v-if="item.temperatureMax / 10 < 5" class="bg-emerald-600 flex items-center justify-center w-full">{{ item.temperatureMax / 10 }} °C</span>
                <span v-else-if="item.temperatureMax / 10 < 10" class="bg-green-700 flex items-center justify-center w-full">{{ item.temperatureMax / 10 }} °C</span>
                <span v-else-if="item.temperatureMax / 10 < 15" class="bg-lime-500 text-slate-800 flex items-center justify-center w-full">{{ item.temperatureMax / 10 }} °C</span>
                <span v-else-if="item.temperatureMax / 10 < 20" class="bg-lime-300 text-slate-800 flex items-center justify-center w-full">{{ item.temperatureMax / 10 }} °C</span>
                <span v-else-if="item.temperatureMax / 10 < 25" class="bg-yellow-500 text-slate-800 flex items-center justify-center w-full">{{ item.temperatureMax / 10 }} °C</span>
                <span v-else-if="item.temperatureMax / 10 < 30" class="bg-orange-500 flex items-center justify-center w-full">{{ item.temperatureMax / 10 }} °C</span>
                <span v-else class="bg-red-700 flex items-center justify-center">{{ item.temperatureMax / 10 }} °C</span>
            </div>


            <div v-for="(item, index) in forecast.days" :key="index" class="flex w-full">
                <span v-if="item.temperatureMin / 10 < -5" class="bg-sky-700 flex items-center justify-center rounded-b-md w-full">{{ item.temperatureMin / 10 }} °C</span>
                <span v-if="item.temperatureMin / 10 < 0" class="bg-cyan-500 text-slate-800 flex items-center justify-center rounded-b-md w-full">{{ item.temperatureMin / 10 }} °C</span>
                <span v-if="item.temperatureMin / 10 < 5" class="bg-emerald-600 flex items-center justify-center rounded-b-md w-full">{{ item.temperatureMin / 10 }} °C</span>
                <span v-else-if="item.temperatureMin / 10 < 10" class="bg-green-700 flex items-center justify-center rounded-b-md w-full">{{ item.temperatureMin / 10 }} °C</span>
                <span v-else-if="item.temperatureMin / 10 < 15" class="bg-lime-500 text-slate-800 flex items-center justify-center rounded-b-md w-full">{{ item.temperatureMin / 10 }} °C</span>
                <span v-else-if="item.temperatureMin / 10 < 20" class="bg-lime-300 text-slate-800 flex items-center justify-center rounded-b-md w-full">{{ item.temperatureMin / 10 }} °C</span>
                <span v-else-if="item.temperatureMin / 10 < 25" class="bg-yellow-500 text-slate-800 flex items-center justify-center rounded-b-md w-full">{{ item.temperatureMin / 10 }} °C</span>
                <span v-else-if="item.temperatureMin / 10 < 30" class="bg-orange-500 flex items-center justify-center rounded-b-md w-full">{{ item.temperatureMin / 10 }} °C</span>
                <span v-else class="bg-red-700 flex items-center justify-center rounded-b-md w-full">{{ item.temperatureMin / 10 }} °C</span>
            </div>
            
        </div>
    </div>
</template>

<script setup lang="ts">
    import { ref } from 'vue'
    import axios from 'axios'

    const props = defineProps({
        id: String
    })

    const forecastDates: any = ref([])
    const forecast = ref()

    const getWeatherForecast = () => {
        const url = '/api/weather/forecast?station=' + props.id
        axios
            .get(url)
            .then((response: any) => {
                forecast.value = response.data
                for (let i = 0; i < response.data.days.length; i++) {
                    const date = response.data.days[i].dayDate
                    const dateSplit = date.split('-')
                    forecastDates.value.push(dateSplit[2] + '.' + dateSplit[1] + '.')
                }
            })
            .catch(() => {
                forecast.value = []
            })
    }

    getWeatherForecast()
    setInterval(() => {
        getWeatherForecast()
    }, 300000)
</script>

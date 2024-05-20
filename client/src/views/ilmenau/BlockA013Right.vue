<template>
    <div class="lg:fixed lg:w-full flex inset-y-0 flex-col">
        <IL_HeaderRight />

        <div class="p-8 grid grid-rows-12 sm:grid-cols-3 grid-cols-1 gap-6 h-full overflow-x-auto">
            <div class="h-full row-span-8 row-start-1 col-start-1 border border-zinc-700 rounded-lg grid-box">
                <WeatherForecast id="N5346" />
            </div>

            <div class="h-full row-span-4 row-start-9 col-start-1 border border-zinc-700 rounded-lg grid-box">
                <WarningsNina id="160700000000" />
            </div>

            <div class="h-full row-span-8 row-start-1 col-start-2 border border-zinc-700 rounded-lg grid-box">
                <EventCalendar />
            </div>

            <div class="h-full row-span-4 row-start-9 col-start-2 border border-zinc-700 rounded-lg overflow-hidden grid-box">
                <BcCafe />
            </div>

            <div class="h-full row-span-6 row-start-1 col-start-3 border border-zinc-700 rounded-lg overflow-hidden grid-box">
                <IL_BusDepartures_Mensa />
            </div>

            <div class="h-full row-span-6 row-start-7 col-start-3 border border-zinc-700 rounded-lg overflow-hidden grid-box">
                <DeparturesTrain id="8010178" />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
    import { ref } from 'vue'
    import IL_HeaderRight from '@/components/IL_HeaderRight.vue'
    import DeparturesTrain from '@/components/DeparturesTrain.vue'
    import IL_BusDepartures_Mensa from '@/components/IL_BusDepartures_Mensa.vue'
    import BcCafe from '@/components/BcCafe.vue'
    import WeatherForecast from '@/components/WeatherForecast.vue'
    import WarningsNina from '@/components/WarningsNina.vue'
    import EventCalendar from '@/components/EventCalendar.vue'

    const hours = ref()

    const getCurrentDateTime = () => {
        const date = new Date();
        hours.value = date.getHours();
    }
    
    getCurrentDateTime()
    setInterval(() => { 
        getCurrentDateTime()
    }, 1000)
</script>

<style scoped>
.grid-box {
  display: grid;
  grid-template-rows: auto 1fr;
}
</style>

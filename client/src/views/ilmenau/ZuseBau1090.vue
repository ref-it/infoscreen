<template>
    <div class="lg:fixed lg:w-full flex inset-y-0 flex-col">
        <IL_Header_IA />

        <div class="p-8 grid grid-rows-12 sm:grid-cols-4 grid-cols-1 gap-6 h-full overflow-x-auto">

            <div class="h-full row-span-12 row-start-1 col-start-1 border border-zinc-700 rounded-lg grid-box">
                <CanteenMenu id="46"/>
            </div>

            <div class="h-full row-span-8 row-start-1 col-start-2 border border-zinc-700 rounded-lg grid-box">
                <CanteenMenu id="53" />
            </div>

            <div class="h-full row-span-4 row-start-9 col-start-2 border border-red-800 bg-red-800 rounded-lg overflow-hidden grid-box">
                <IL_FSR_IA />
            </div>

            <!--<div class="h-full row-span-12 row-start-1 col-start-3 border border-zinc-700 rounded-lg grid-box">
                <CanteenMenu id="55" />
            </div>-->

            <div class="h-full row-span-6 row-start-1 col-start-3 border border-zinc-700 rounded-lg grid-box">
                <WeatherForecast id="N5346" />
            </div>

            <div class="h-full row-span-6 row-start-7 col-start-3 border border-zinc-700 rounded-lg overflow-hidden grid-box">
                <IL_BusDepartures_Kirchhoff />
            </div>

            <div class="h-full row-span-8 row-start-1 col-start-4 border border-zinc-700 rounded-lg grid-box">
                <EventCalendar />
            </div>

            <div class="h-full row-span-4 row-start-9 col-start-4 border border-zinc-700 rounded-lg overflow-hidden grid-box">
                <BcCafe />
            </div>

        </div>

        <!--<PageFooter />-->
    </div>
</template>

<script setup lang="ts">
    import { ref } from 'vue'
    import IL_Header_IA from '@/components/IL_Header_IA.vue'
    import CanteenMenu from '@/components/CanteenMenu.vue'
    import IL_BusDepartures_Kirchhoff from '@/components/IL_BusDepartures_Kirchhoff.vue'
    import BcCafe from '@/components/BcCafe.vue'
    import WeatherForecast from '@/components/WeatherForecast.vue'
    import EventCalendar from '@/components/EventCalendar.vue'
    import IL_FSR_IA from '@/components/IL_FSR_IA.vue'

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

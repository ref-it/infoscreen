<template>
    <div class="lg:fixed lg:w-full flex inset-y-0 flex-col">
        <IL_Header />

        <div class="p-8 grid grid-rows-12 sm:grid-cols-4 grid-cols-1 gap-6 h-full overflow-x-auto">
            <div v-if="hours < 14" class="h-full row-span-12 row-start-1 col-start-1 border border-slate-700 rounded-lg grid-box">
                <CanteenMenu id="46" />
            </div>

            <div v-if="hours < 14" class="h-full row-span-12 row-start-1 col-start-2 border border-slate-700 rounded-lg grid-box">
                <CanteenMenu id="53" />
            </div>

            <div v-if="hours >= 14" class="h-full row-span-12 row-start-1 col-start-1 border border-slate-700 rounded-lg grid-box">
                <CanteenMenu id="53" />
            </div>

            <!--<div v-if="hours < 14" class="h-full row-span-12 row-start-1 col-start-3 border border-slate-700 rounded-lg grid-box">
                <CanteenMenu id="55"/>
            </div>-->

            <!--<div v-if="hours >= 14" class="h-full row-span-12 col-span-2 row-start-1 col-start-2 border border-slate-700 rounded-lg grid-box">
                Vereinswerbung
            </div>-->

            <div class="h-full row-span-5 row-start-1 col-start-3 border border-slate-700 rounded-lg grid-box">
                <WeatherForecast id="N5346" />
            </div>

            <div class="h-full row-span-4 row-start-6 col-start-3 border border-slate-700 rounded-lg grid-box">
                <WebcamIlmenau />
            </div>

            <div class="h-full row-span-6 row-start-1 col-start-4 border border-slate-700 rounded-lg grid-box">
                <IL_Events />
            </div>

            <!--<div class="h-full row-span-1 row-start-4 col-start-3 border border-slate-700 rounded-lg grid-box">
                <IL_Konsul />
            </div>-->

            <!--<div class="h-full row-span-2 row-start-1 col-start-4 border border-slate-700 rounded-lg overflow-hidden grid-box">
                <IL_BusDepartures_Kirchhoff />
            </div>-->

            <div class="h-full row-span-6 row-start-7 col-start-4 border border-slate-700 rounded-lg overflow-hidden grid-box">
                <DeparturesTrain id="8010178" />
            </div>

            <!--<div class="h-full row-span-2 row-start-3 col-start-4 border border-slate-700 rounded-lg overflow-hidden grid-box">
                <DeparturesTrain id="8010101" />
            </div>-->
        </div>

        <PageFooter />
    </div>
</template>

<script setup lang="ts">
    import { ref } from 'vue'
    import IL_Header from '@/components/IL_Header.vue'
    import PageFooter from '@/components/PageFooter.vue'
    import CanteenMenu from '@/components/CanteenMenu.vue'
    import DeparturesTrain from '@/components/DeparturesTrain.vue'
    import IL_BusDepartures_Kirchhoff from '@/components/IL_BusDepartures_Kirchhoff.vue'
    import IL_Events from '@/components/IL_Events.vue'
    import IL_Konsul from '@/components/IL_Konsul.vue'
    import WebcamIlmenau from '@/components/WebcamIlmenau.vue'
    import WeatherForecast from '@/components/WeatherForecast.vue'

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
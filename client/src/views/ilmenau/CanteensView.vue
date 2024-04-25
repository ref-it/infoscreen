<template>
    <div class="lg:fixed lg:w-full flex inset-y-0 flex-col">
        <IL_Header />

        <div class="p-8 grid grid-rows-4 grid-cols-4 grid-cols-1 gap-6 h-full overflow-x-auto">
            <div class="h-full row-span-4 row-start-1 col-start-1 border border-slate-700 rounded-lg grid-box">
                <CanteenMenu id="46"/>
            </div>

            <div class="h-full row-span-4 row-start-1 col-start-2 border border-slate-700 rounded-lg grid-box">
                <CanteenMenu id="53" />
            </div>

            <div class="h-full row-span-4 row-start-1 col-start-3 border border-slate-700 rounded-lg grid-box">
                <CanteenMenu id="55" />
            </div>

            <div class="h-full row-span-4 row-start-1 col-start-4 border border-slate-700 rounded-lg grid-box">
                <CanteenMenu id="57" />
            </div>
        </div>

        <PageFooter />
    </div>
</template>

<script setup lang="ts">
import IL_Header from '@/components/IL_Header.vue'
import CanteenMenu from '@/components/CanteenMenu.vue'
import PageFooter from '@/components/PageFooter.vue'
</script>

<style scoped>
.grid-box {
  display: grid;
  grid-template-rows: auto 1fr;
}
</style>
<template>
    <div class="sticky top-0 z-40 flex h-24 shrink-0 items-center gap-x-4 border-b px-4 shadow-sm shadow-black/20 sm:gap-x-6 px-8 bg-slate-800 text-white border-slate-900">
        <div class="flex flex-1 gap-x-4 self-stretch lg:gap-x-6 text-4xl">
            <div class="flex items-center gap-x-4 lg:gap-x-6">
                <router-link to="/">
                    <img src="@/assets/img/stura-logo-breit-white.svg" class="h-16">
                </router-link>
            </div>
            <div class="flex items-center gap-x-4 lg:gap-x-6 ml-auto">
                <DateTime />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
    import DateTime from '@/components/DateTime.vue'
</script>
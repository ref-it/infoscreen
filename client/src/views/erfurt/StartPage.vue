<template>
    <div class="lg:fixed lg:w-full flex inset-y-0 flex-col">
        <IL_Header />

        <div class="p-8 grid grid-cols-4 grid-cols-1 gap-6 overflow-x-auto">
            <div v-for="place in screens" :key="place.title" class="h-full border border-zinc-700 rounded-lg relative p-6 focus-within:ring-2 focus-within:ring-inset focus-within:ring-yellow-400 border border-black/5 transition ease-in-out duration-300 rounded-lg shadow hover:bg-gray-100 bg-gray-50 dark:bg-zinc-700 dark:hover:bg-zinc-600">
                <div>
                    <h3 class="text-2xl font-bold leading-6 text-gray-900 dark:text-white">
                        <router-link :to="place.href" class="focus:outline-none">
                            <!-- Extend touch target to entire panel -->
                            <span class="absolute inset-0" aria-hidden="true"></span>
                            {{ place.title }}
                        </router-link>
                    </h3>
                </div>
                <span class="pointer-events-none absolute right-6 top-6 text-gray-300 group-hover:text-gray-400" aria-hidden="true">
                    <svg class="h-6 w-6" fill="currentColor" viewBox="0 0 24 24">
                        <path d="M20 4h1a1 1 0 00-1-1v1zm-1 12a1 1 0 102 0h-2zM8 3a1 1 0 000 2V3zM3.293 19.293a1 1 0 101.414 1.414l-1.414-1.414zM19 4v12h2V4h-2zm1-1H8v2h12V3zm-.707.293l-16 16 1.414 1.414 16-16-1.414-1.414z" />
                    </svg>
                </span>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import IL_Header from '@/components/IL_Header.vue'

const screens = [
    {
        title: 'Mensen & Kantinen',
        href: '/ef/canteens',
    }
]
</script>
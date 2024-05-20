<template>
    <div class="sticky top-0 z-40 flex h-24 shrink-0 items-center gap-x-4 border-b px-4 shadow-sm shadow-black/20 sm:gap-x-6 px-8 bg-zinc-800 text-white border-zinc-900">
        <div class="flex flex-1 gap-x-4 self-stretch lg:gap-x-6 text-4xl">
            <div class="flex items-center gap-x-4 lg:gap-x-6">
                <router-link to="/">
                    <img src="@/assets/img/fsr-ia-logo-breit-white.svg" class="h-11">
                </router-link>
            </div>
            <div class="flex items-center gap-x-4 lg:gap-x-6 ml-auto">
                <DateTime />
            </div>
            <div class="flex items-center">
                <div class="h-8 w-px bg-zinc-700 mr-4 lg:mr-6" aria-hidden="true"></div>
                <button @click="infoOpen = true">
                    <InformationCircleIcon class="h-11 w-11 text-white" aria-hidden="true" />
                </button>
            </div>
        </div>
    </div>

    <TransitionRoot as="template" :show="infoOpen">
        <Dialog as="div" class="relative z-10" @close="infoOpen = false">
            <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0" enter-to="opacity-100" leave="ease-in duration-200" leave-from="opacity-100" leave-to="opacity-0">
                <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
            </TransitionChild>

            <div class="fixed inset-0 z-10 w-screen overflow-y-auto">
                <div class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
                    <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95" enter-to="opacity-100 translate-y-0 sm:scale-100" leave="ease-in duration-200" leave-from="opacity-100 translate-y-0 sm:scale-100" leave-to="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95">
                        <DialogPanel class="relative transform overflow-hidden rounded-lg bg-zinc-800 px-4 pb-4 pt-5 text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-xl sm:p-6">
                        <div class="absolute right-0 top-0 hidden pr-4 pt-4 sm:block">
                            <button type="button" class="rounded-md bg-zinc-800 text-white hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2" @click="infoOpen = false">
                            <span class="sr-only">Close</span>
                            <XMarkIcon class="h-6 w-6" aria-hidden="true" />
                            </button>
                        </div>
                        <div class="sm:flex sm:items-start">
                            <div class="mx-auto flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-full bg-yellow-100 sm:mx-0 sm:h-10 sm:w-10">
                            <InformationCircleIcon class="h-6 w-6 text-yellow-600" aria-hidden="true" />
                            </div>
                            <div class="mt-3 text-center sm:ml-4 sm:mt-0 sm:text-left">
                            <DialogTitle as="h3" class="text-base font-semibold leading-6 text-white">StuRa-InfoScreen</DialogTitle>
                            <div class="mt-2">
                                <InfoAbout />
                            </div>
                            </div>
                        </div>
                        </DialogPanel>
                    </TransitionChild>
                </div>
            </div>
        </Dialog>
    </TransitionRoot>
</template>

<script setup lang="ts">
    import { ref } from 'vue'
    import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot } from '@headlessui/vue'
    import { XMarkIcon, InformationCircleIcon } from '@heroicons/vue/24/outline'
    import DateTime from '@/components/DateTime.vue'
    import InfoAbout from '@/components/InfoAbout.vue'

    const infoOpen = ref(false)

</script>
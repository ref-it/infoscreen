<template>
    {{ day }}
    <div class="h-8 w-px bg-slate-700" aria-hidden="true"></div>
    {{ time }}
</template>

<script setup lang="ts">
    import { ref } from 'vue'

    let time = ref()
    let day = ref()
    let weekday = ref()

    const getCurrentDateTime = () => {
        const date = new Date();
        weekday.value = date.getDay();
        day.value = Intl.DateTimeFormat('de-DE', {
            weekday: 'long',
            day: '2-digit',
            month: '2-digit',
            year: 'numeric'
        }).format()

        time.value = Intl.DateTimeFormat('de-DE', {
            hour: 'numeric',
            minute: 'numeric',
            second: 'numeric'
        }).format()
    }

    getCurrentDateTime()
    setInterval(() => { 
        getCurrentDateTime()
    }, 1000)
</script>
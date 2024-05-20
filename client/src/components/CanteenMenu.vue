<template>
    <div class="bg-black/20 text-xl font-bold p-4 border-b border-zinc-700 pl-4 pr-2 py-2 rounded-t-lg grid header-with-badge">
        <span>{{ canteenName }}</span>
        <div class="-mt-02">
            <span v-if="props.id == 53 && hours < 17" class="inline-flex items-center rounded-md bg-emerald-400/10 ml-auto px-2 py-1 text-sm text-emerald-400 ring-1 ring-inset ring-emerald-400/30">Mittag</span>
            <span v-else-if="props.id == 53 && hours >= 17" class="inline-flex items-center rounded-md bg-emerald-400/10 ml-auto px-2 py-1 text-sm text-emerald-400 ring-1 ring-inset ring-emerald-400/30">Abend</span>
        </div>
    </div>
    <div class="h-full overflow-y-auto flex flex-col">
        <div v-if="mealsDinner != null && mealsDinner.length > 0 && hours >= 17">
            <div v-for="(meal, index) in mealsDinner" :key="index" class="border-b border-zinc-700 grid meal-item">
                <div class="px-4 py-2 pb-3 border-l-4" :class="(meal.isVegetarian && !meal.isVegan)?'border-lime-400':'' || (meal.isVegan)?'border-green-700':'' || (!meal.isVegetarian && !meal.isVegan)?'border-zinc-800':''">
                    <div class="grid meal-item">
                        <div class="pr-3 border-r border-zinc-700">
                            <div class="text-xl mb-2 font-bold">{{ meal.name }}</div>
                            <div v-if="meal.additives.length > 0" class="text-sm"><b>Z:</b> <span>{{ meal.additives[0] }}</span><span v-for="(additive, index) in meal.additives.slice(1)" :key="index">, {{ additive }}</span></div>
                            <div v-if="meal.allergens.length > 0" class="text-sm"><b>A:</b> <span>{{ meal.allergens[0] }}</span><span v-for="(allergen, index) in meal.allergens.slice(1)" :key="index">, {{ allergen }}</span></div>
                        </div>
                        <div class="text-right text-base pl-3">
                            <b>{{ currencyFormatter.format(meal.prices.students) }}</b><br />
                            {{ currencyFormatter.format(meal.prices.employees) }}<br />
                            {{ currencyFormatter.format(meal.prices.guests) }}
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div v-else-if="mealsLunch != null && mealsLunch.length > 0 && mergeProps.id != 53 || mealsLunch != null && mealsLunch.length > 0 && mergeProps.id == 53 && hours < 17">
            <div v-for="(meal, index) in mealsLunch" :key="index" class="border-b border-zinc-700 grid meal-item">
                <div class="px-4 py-2 pb-3 border-l-4" :class="(meal.isVegetarian && !meal.isVegan)?'border-lime-400':'' || (meal.isVegan)?'border-green-700':'' || (!meal.isVegetarian && !meal.isVegan)?'border-zinc-800':''">
                    <div class="grid meal-item">
                        <div class="pr-3 border-r border-zinc-700">
                            <div class="text-xl font-bold">{{ meal.name }}</div>
                            <div v-if="meal.additives.length > 0" class="text-sm"><b>Z:</b> <span>{{ meal.additives[0] }}</span><span v-for="(additive, index) in meal.additives.slice(1)" :key="index">, {{ additive }}</span></div>
                            <div v-if="meal.allergens.length > 0" class="text-sm"><b>A:</b> <span>{{ meal.allergens[0] }}</span><span v-for="(allergen, index) in meal.allergens.slice(1)" :key="index">, {{ allergen }}</span></div>
                        </div>
                        <div class="text-right text-base pl-3">
                            <b>{{ currencyFormatter.format(meal.prices.students) }}</b><br />
                            {{ currencyFormatter.format(meal.prices.employees) }}<br />
                            {{ currencyFormatter.format(meal.prices.guests) }}
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div v-else class="px-4 py-2 text-xl m-auto text-center">
            <p>Heute leider kein Essen in dieser Einrichtung.</p>
        </div>
    </div>
</template>
  
<script setup lang="ts">
    import { mergeProps, ref } from 'vue'
    import axios from 'axios'

    const props = defineProps({
        id: String
    })

    const weekday = ref()
    const hours = ref()

    const getCurrentDateTime = () => {
        const date = new Date();
        weekday.value = date.getDay();
        hours.value = date.getHours() + 1;
    }

    const canteenName: any = ref()
    const mealsLunch: any = ref([])
    const mealsDinner: any = ref([])
  
    const currencyFormatter = new Intl.NumberFormat('de-DE', {
        style: 'currency',
        currency: 'EUR',
    })

    const getMensaMeals = () => {
        const url = '/api/menu?canteen=' + props.id
        axios
            .get(url)
            .then((response: any) => {
                console.log(response.data)
                canteenName.value = response.data.canteenName
                mealsLunch.value = response.data.lunch
                mealsDinner.value = response.data.dinner
            })
            .catch(() => {
                mealsLunch.value = []
                mealsDinner.value = []
            })
    }
    
    getCurrentDateTime()
    setInterval(() => { 
        getCurrentDateTime()
    }, 1000)

    getMensaMeals()
    setInterval(() => {
        getMensaMeals()
    }, 300000)
</script>

<style scoped>
    .grid-box {
        display: grid;
        grid-template-rows: auto 1fr;
    }

    .grid-box2 {
        display: grid;
        grid-template-rows: 1fr auto;
    }

    .meal-item {
        grid-template-columns: 1fr auto;
    }
</style>

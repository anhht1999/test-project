<template>
    <div>
        <MainSilder/>
        <MainProduct/>
    </div>
</template>

<script>

import MainSilder from "./MainSilder.vue";
import MainProduct from "./MainProduct.vue";
export default {
    name:"HomePage",
    components: {
        MainSilder,
        MainProduct,
    }
}
</script>
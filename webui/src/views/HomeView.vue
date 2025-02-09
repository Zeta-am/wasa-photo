<template>
	<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
	  <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		<h1 class="h2">Home page</h1>
	  </div>
	  <ErrorMsg v-if="errormsg" :msg="errormsg"/>
	</main>
  </template>

<script>
export default {
  data() {
    return {
      errormsg: null
    }
  }
}
</script>

<style>
</style>

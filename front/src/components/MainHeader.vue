<template>
  <header class="header shop">
		<!-- Topbar -->
		<!-- End Topbar -->
		<div class="middle-inner">
			<div class="container">
				<div class="row">
					<div class="col-lg-2 col-md-2 col-12">
						<!-- Logo -->
						<div class="logo">
							<a href="index.html"><img src="images/logo.png" alt="logo"></a>
						</div>
						<!--/ End Logo -->
						<!-- Search Form -->
						<div class="search-top">
							<div class="top-search"><a href="#0"><i class="ti-search"></i></a></div>
							<!-- Search Form -->
							<div class="search-top">
								<form class="search-form">
									<input type="text" placeholder="Search here..." name="search">
									<button value="search" type="submit"><i class="ti-search"></i></button>
								</form>
							</div>
							<!--/ End Search Form -->
						</div>
						<!--/ End Search Form -->
						<div class="mobile-nav"></div>
					</div>
					<div class="col-lg-8 col-md-7 col-12">
						<!-- <div class="search-bar-top">
							<div class="search-bar">
								<form>
									<input name="search" placeholder="Search Products Here....." type="search">
									<button class="btnn"><i class="ti-search"></i></button>
								</form>
							</div>
						</div> -->
					</div>
					<div class="col-lg-2 col-md-3 col-12">
						<div class="right-bar">
							<!-- Search Form -->
							
							<div class="sinlge-bar">
								<router-link to="/login" class="single-icon"><i class="fa fa-user-circle-o"></i></router-link>
							</div>
							<div class="sinlge-bar shopping">
								<router-link to="/cart" class="single-icon"><i class="ti-bag"></i> <span class="total-count">{{totalItems}}</span></router-link>
								<!-- Shopping Item -->
								<div class="shopping-item">
									<div class="dropdown-cart-header">
										<span>{{totalItems}}</span>
										<router-link to="/cart">View Cart</router-link>
									</div>
									<ul class="shopping-list" v-for="cart in carts" :key="cart.ID">
										<li>
											<!-- <a href="#" class="remove" title="Remove this item"><i class="fa fa-remove"></i></a> -->
											<a class="cart-img" href="#"><img :src="'http://localhost:3000' + cart.Url" alt="#"></a>
											<h4><a href="#">{{cart.Name}}</a></h4>
											<p class="quantity">{{cart.quantity}} - <span class="amount">${{cart.Price}}</span></p>
										</li>
									</ul>
									<div class="bottom">
										<div class="total">
											<span>Total</span>
											<span class="total-amount">${{subTotal}}</span>
										</div>
										<a href="#" class="btn animate">Checkout</a>
									</div>
								</div>
								<!--/ End Shopping Item -->
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
		<!-- Header Inner -->
		<div class="header-inner">
			<div class="container">
				<div class="cat-nav-head">
					<div class="row">
						<div class="col-lg-9 col-12">
							<div class="menu-area">
								<!-- Main Menu -->
								<nav class="navbar navbar-expand-lg">
									<div class="navbar-collapse">	
										<div class="nav-inner">	
											<ul class="nav main-menu menu navbar-nav">
													<li class="active"><router-link to="/">Home</router-link></li>
													<li><router-link to="/product">Product</router-link></li>												
													
													<li><router-link to="/about">Contact Us</router-link></li>
												</ul>
										</div>
									</div>
								</nav>
								<!--/ End Main Menu -->	
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
		<!--/ End Header Inner -->
	</header>
</template>

<script>
import { mapState, mapGetters} from "vuex";

export default {

  name: 'MainHeader',

  computed: {
    ...mapState("cart", ["carts"]),
    ...mapGetters("cart", ["totalItems"]),
    ...mapGetters("cart", ["subTotal"]),
  },
 
}
</script>

<style>
</style>

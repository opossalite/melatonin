<script lang="ts">
    import { fly } from "svelte/transition";
    import Header from "./Header.svelte";
    import MainWindow from "./MainWindow.svelte";
    import Footer from "./Footer.svelte";



//<div id="wrapper">
//    <Header name={"a"}/>
//    <MainWindow></MainWindow>
//    <Footer></Footer>
//</div>



</script>

<div id="wrapper">
    <div id="top"><Header name="a"/></div>
    <div id="middle"><MainWindow/></div>
    <div id="bottom"><Footer/></div>
</div>





<style>


:global(html,body),#wrapper {
    height:100%;
    padding:0;
    margin:0;
}
#wrapper {
    position: relative;
}

#top, #middle, #bottom {
    position:absolute;
}

#top {
    height:50px;
    width:100%;
    background:grey;
}
#middle {
    top:50px;
    bottom:50px;
    width:100%;
    background:black;
    color:white;
}
#bottom {
    bottom:0;
    height:50px;
    width:100%;
    background:grey;
}


.error {
    color: red;
}

.horiz {
    display: inline-block;
}
</style>





<template>
<div id='container'>
  <span></span>
  <span></span>
  <span></span>
  <span></span>
  <span></span>
  <span></span>
  <span></span>
</div>
</template>
<<script>
export default {
  name: 'Waves',
}
</script>

<style>
#container{
  position:relative;
}

#container span{
  width:5px;
  height: 5px;
  bottom:0px;
  position:absolute;
  background: #e2c044;
  -webkit-animation: bodong 0.5s infinite  ease;
}

#container span:first-child{
  left:0px;
  -webkit-animation-delay:.3s;
}

#container span:nth-child(2){
  left:7px; 
  -webkit-animation-delay:.4s;
}
#container span:nth-child(3){
  left:14px;
  -webkit-animation-delay:.6s;
}
#container span:nth-child(4){
  left:21px;
  -webkit-animation-delay:.8s;
}
#container span:nth-child(5){
  left:28px;
  -webkit-animation-delay:1s;
}

@-webkit-keyframes bodong{  
  0%{height:5px; background:bule;}  
  30%{height:15px; background:bule;}  
  60%{height:20px; background:bule;}  
  80%{height:15px; background:bule;}  
  100%{height:5px; background:bule;}  
}
</style>

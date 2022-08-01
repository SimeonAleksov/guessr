<template>
  <div>
  <div class="flex flex-col items-center z-40">
    <div class="mt-12 flex space-x-4 z-1">
      <div class="dashboard__player dark:bg-dark text-primary px-4 uppercase flex flex-col justify-center">
        <h1 class="text-center text-secondary font-semibold">1209</h1>
        <img class="w-24 h-24" src='https://avataaars.io/?avatarStyle=Transparent&topType=ShortHairShortFlat&accessoriesType=Blank&hairColor=Blue&facialHairType=BeardMedium&facialHairColor=Blonde&clotheType=Hoodie&clotheColor=Blue03&eyeType=Dizzy&eyebrowType=UpDown&mouthType=Disbelief&skinColor=Pale'
        />
        <h1 class="text-center"> Player 3 </h1>
      </div>
      <div class="dashboard__player dark:bg-dark text-primary px-4 uppercase flex flex-col">
        <h1 class="text-center text-secondary font-semibold">1209</h1>
        <img class="w-24 h-24" src='https://avataaars.io/?avatarStyle=Transparent&topType=ShortHairDreads02&accessoriesType=Prescription02&hatColor=Gray01&hairColor=BrownDark&facialHairType=BeardMedium&facialHairColor=BrownDark&clotheType=CollarSweater&clotheColor=PastelOrange&eyeType=WinkWacky&eyebrowType=UnibrowNatural&mouthType=Disbelief&skinColor=DarkBrown'
        />
        <h1 class="text-center"> Player 5 </h1>
      </div>
      <div class="dashboard__player dark:bg-dark text-primary px-4 uppercase flex flex-col">
        <h1 class="text-center text-secondary font-semibold">529</h1>
        <img class="w-24 h-24" src='https://avataaars.io/?avatarStyle=Transparent&topType=LongHairStraight2&accessoriesType=Sunglasses&hairColor=Red&facialHairType=BeardMedium&facialHairColor=BlondeGolden&clotheType=GraphicShirt&clotheColor=Blue02&graphicType=Selena&eyeType=Surprised&eyebrowType=SadConcernedNatural&mouthType=Vomit&skinColor=Light'/>
        <h1 class="text-center"> Player 1 </h1>
      </div>
      <div class="dashboard__player dark:bg-dark text-primary px-4 uppercase flex flex-col">
        <h1 class="text-center text-secondary font-semibold">123</h1>
        <img class="w-24 h-24" src="https://avataaars.io/?avatarStyle=Transparent&topType=NoHair&accessoriesType=Sunglasses&facialHairType=Blank&clotheType=GraphicShirt&clotheColor=Heather&graphicType=Skull&eyeType=Default&eyebrowType=RaisedExcited&mouthType=ScreamOpen&skinColor=Black"/>
        <h1 class="text-center"> Player 5 </h1>
      </div>
      <div class="dashboard__player dark:bg-dark text-primary px-4 uppercase flex flex-col">
        <h1 class="text-center text-secondary font-semibold">48</h1>
        <img class="w-24 h-24" src="https://avataaars.io/?avatarStyle=Transparent&topType=LongHairBigHair&accessoriesType=Kurt&hairColor=Black&facialHairType=Blank&clotheType=ShirtCrewNeck&clotheColor=Red&eyeType=Surprised&eyebrowType=SadConcernedNatural&mouthType=ScreamOpen&skinColor=Pale"/>
        <h1 class="text-center"> Player 2 </h1>
      </div>
    </div>
    <div class="mt-32 grid grid-rows-2 grid-flow-col gap-4 w-full h-2/6">
      <TriviaButton class="w-full">
        Seks I Droga
      </TriviaButton>
      <TriviaButton class="w-full">
        Pizda Materina
      </TriviaButton>
      <TriviaButton class="w-full">
        Samac
      </TriviaButton>
      <TriviaButton class="w-full">
        Šef
      </TriviaButton>
    </div>
    <Vinyl class="mt-24"/>
    <video allow="autoplay" ref="video" width="100%" controls playsinline="" hidden autoplay>
    </video>
  </div>
  </div>
</template>
<script>
import Hls from 'hls.js'
import Button from '../../components/shared/Button.vue'
import TriviaButton from '../../components/shared/TriviaButton.vue'
import Waves from '../../components/Waves.vue'
import Vinyl from '../../components/Vinyl.vue'
export default {
  name: 'trivia',
  middleware: 'auth',
  components: {
      Button,
      TriviaButton,
      Waves,
      Vinyl,
  },
  mounted() {
    let hls = new Hls();
    let video = this.$refs["video"];
    hls.attachMedia(video);
    hls.on(Hls.Events.MEDIA_ATTACHED, function () {
      console.log('video and hls.js are now bound together !');
      hls.loadSource('http://localhost:8000/songs/soad/lost-in-hollywood/outputlist.m3u8');
      hls.on(Hls.Events.MANIFEST_PARSED, function (event, data) {
        console.log(
          'manifest loaded, found ' + data.levels.length + ' quality level'
        );

      video.play()
      });
    });
    console.log(video)
  },
}
</script>
<<style lang="scss">
.dashboard {
  &__player {
    box-shadow: rgba(50, 50, 93, 0.25) 0px 2px 5px -1px, rgba(0, 0, 0, 0.3) 0px 1px 3px -1px;
  }
}

.vinyl {
}
</style>

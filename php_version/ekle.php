
<?php

if(isset($_POST["kaydet"]))
{
  $icerik1 = $_POST['icerik'];
  $icerik2 = implode($icerik1,",");
    include "baglan.php";
    $sql = "INSERT INTO `yardimlar4` (`id`, `il`, `ilce`, `tel`, `detay`, `date`, `icerik`, `adres`) VALUES (NULL, ?, ?, ?, ?, current_timestamp(), '$icerik2', ?);";
    $dizi =[
        $_POST["il"],
        $_POST["ilce"],
        $_POST["tel"],
        $_POST["detay"],
        $_POST["adres"],
    ];
    $sth = $baglan->prepare($sql);
   $sonuc = $sth->execute($dizi);
   header("Location:index.php");
}

?>
<!DOCTYPE html>
<html lang="tr">
<body style="background-color:#181a1b;">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Yardım Toplanma Bölgeleri</title>
    <!-- CSS only -->
    <script type="text/javascript" src="//code.jquery.com/jquery-3.3.1.min.js"></script>
<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.1/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-+0n0xVW2eSR5OomGNYDnhzAbDsOXxcvSN1TPprVMTNDbiYZCxYbOOl7+AMvyTG2x" crossorigin="anonymous">
</head>
<body>
  <style>
    .form-label {
      color: white;
    }
    label {
      color: white;
    }
</style>
      <header>
      <div class="p-5 bg-primary text-white text-center">
        <h2>Yardım Bulunan Bölgeler</h2>
        <h6>Yardım ulaşan bölgelerdeki vatandaşların yardımlardan haberdar olabilmesi için, lütfen teyitli noktalardaki yardımları depremzedeler ile paylaşınız.</h6>
      </div>

      <nav class="navbar navbar-expand-sm bg-dark navbar-dark">
        <div class="container-fluid">
          <ul class="navbar-nav mx-auto">
            <li class="nav-item">
              <a class="nav-link" href="index.php">Yardımlar</a>
            </li>
            <li class="nav-item">
              <a class="nav-link active" href="ekle.php">Yardım Ekle</a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="kizilay.php">Kızılay</a>
            </li>
          </ul>
        </div>
      </nav>
    </header>
    <main>
    <div class="container mt-3">
      <h2 class="text-white">Yardım Ekle</h2>
      <form name="afadd" action="" method="post">
        <div class="mb-3 mt-3">
          <label for="il">İl*:</label>
          <select class="form-select" id="il" name="il">
        <option>İl seçiniz</option><option>Adana</option><option>Adiyaman</option><option>Afyon</option><option>Agri</option><option>Amasya</option><option>Ankara</option><option>Antalya</option><option>Artvin</option><option>Aydin</option><option>Balikesir</option><option>Bilecik</option><option>Bingol</option><option>Bitlis</option><option>Bolu</option><option>Burdur</option><option>Bursa</option><option>Canakkale</option><option>Cankiri</option><option>Corum</option><option>Denizli</option><option>Diyarbakir</option><option>Edirne</option><option>Elazig</option><option>Erzincan</option><option>Erzurum</option><option>Eskisehir</option><option>Gaziantep</option><option>Giresun</option><option>Gumushane</option><option>Hakkari</option><option>Hatay</option><option>Isparta</option><option>Mersin</option><option>Istanbul</option><option>Izmir</option><option>Kars</option><option>Kastamonu</option><option>Kayseri</option><option>Kirklareli</option><option>Kirsehir</option><option>Kocaeli</option><option>Konya</option><option>Kutahya</option><option>Malatya</option><option>Manisa</option><option>Kahramanmaras</option><option>Mardin</option><option>Mugla</option><option>Mus</option><option>Nevsehir</option><option>Nigde</option><option>Ordu</option><option>Rize</option><option>Sakarya</option><option>Samsun</option><option>Siirt</option><option>Sinop</option><option>Sivas</option><option>Tekirdag</option><option>Tokat</option><option>Trabzon</option><option>Tunceli</option><option>Sanliurfa</option><option>Usak</option><option>Van</option><option>Yozgat</option><option>Zonguldak</option><option>Aksaray</option><option>Bayburt</option><option>Karaman</option><option>Kirikkale</option><option>Batman</option><option>Sirnak</option><option>Bartin</option><option>Ardahan</option><option>Igdir</option><option>Yalova</option><option>Karabuk</option><option>Kilis</option><option>Osmaniye</option><option>Duzce</option>
          </select>
        </div>
        <div class="mb-3 mt-3">
          <label for="ilce">İlçe*:</label>
          <input type="text" class="form-control" id="ilce" placeholder="İlçe giriniz" name="ilce" required>
        </div>
            <div class="mb-3 mt-3">
          <label for="adres">Açık Adres*:</label>
          <input type="text" class="form-control" id="adres" placeholder="Açık adres giriniz" name="adres" required>
        </div>
        	<div class="mb-3 mt-3">
          <label for="tel">Telefon:</label>
          <input type="tel" class="form-control" id="tel" placeholder="05555555555" pattern="[0-9]{11}" name="tel">
        </div>
            <div class="mb-3 mt-3">
          <label for="detay">Açıklama*:</label>
          <input type="text" class="form-control" id="detay" placeholder="Yardım içeriğinizi detaylandırınız" name="detay" required>
        </div>
                <div class="col-6 text-white">
                  Yardım İçeriği*:</br>
                  <input type="checkbox" id="icerik" name="icerik[]" value="Gıda" /> Gıda</br>
                  <input type="checkbox" id="icerik" name="icerik[]" value="Barınma" /> Barınma</br>
                  <input type="checkbox" id="icerik" name="icerik[]" value="Sağlık" /> Sağlık</br>
                  <input type="checkbox" id="icerik" name="icerik[]" value="Giyim" /> Giyim</br>
                  <input type="checkbox" id="icerik" name="icerik[]" value="Ekipman" /> Ekipman</br>
                </div>
                </br>
        <button type="submit" name="kaydet" class="btn btn-primary">Kaydet</button>
      </form>
    </div>

    </main>
    <footer><div style="font-size: 80%;" class="mt-5 p-4 bg-dark text-white text-center">
  <p>Bu site yardım amaçlı kurulmuş olup, herhangi bir maddi kazanç gözetmemektedir.
Sitede paylaşılan bilgiler depremzedelerin ve yardımsever vatandaşların birbirleri ile kolay şekilde iletişim kurabilmelerini sağlamak adına depolanacaktır.</p>
</div></footer>
</body>
</html>

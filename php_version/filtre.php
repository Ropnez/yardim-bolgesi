
<?php

include "baglan.php";

$sql ="SELECT * FROM yardimlar4 WHERE il = ? ORDER BY id DESC";
$sorgu = $baglan->prepare($sql);
$sorgu->execute([
    $_GET['il']
]);
?>
<!DOCTYPE html>
<html lang="tr">
<body style="background-color:#181a1b;">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=0>
    <title>Yardım Bulunan Bölgeler</title>
    <!-- CSS only -->

<link rel="stylesheet" href="style.css" type="text/css"/>
<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.1/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-+0n0xVW2eSR5OomGNYDnhzAbDsOXxcvSN1TPprVMTNDbiYZCxYbOOl7+AMvyTG2x" crossorigin="anonymous">
</head>
<body>
<style>
@media only screen and (max-width: 600px) {
  .container, .container-lg, .container-md, .container-sm, .container-xl, .container-xxl {
    width: 100%;
    padding-right: var(--bs-gutter-x,.75rem);
    padding-left: var(--bs-gutter-x,.0rem);
    margin-right: auto;
    margin-left: auto;
    font-size: 8px
  }
  @media only screen and (max-width: 600px) {
    .btn.btn-success {
      font-size: 6px
    }
    @media only screen and (max-width: 600px) {
      .btn.btn-secondary {
        font-size: 6px
      }
      @media only screen and (max-width: 600px) {
        .display-1.text-center {
          font-size: 35px
        }
        @media only screen and (max-width: 600px) {
          .btn.btn-primary mt-3 {
            font-size: 12px
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
            <a class="nav-link active" href="index.php">Yardımlar</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="ekle.php">Yardım Ekle</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="kizilay.php">Kızılay</a>
          </li>
        </ul>
      </div>
    </nav>
        <div class="container mt-3 mx-auto">
  <form action="/filtre.php?il=<?=$satir['il']?>">
    <h6 for="il" class="form-label">İl Seçiniz:</h6>
    <select class="form-select" id="il" name="il">
  <option>İl seçiniz</option><option>Adana</option><option>Adiyaman</option><option>Afyon</option><option>Agri</option><option>Amasya</option><option>Ankara</option><option>Antalya</option><option>Artvin</option><option>Aydin</option><option>Balikesir</option><option>Bilecik</option><option>Bingol</option><option>Bitlis</option><option>Bolu</option><option>Burdur</option><option>Bursa</option><option>Canakkale</option><option>Cankiri</option><option>Corum</option><option>Denizli</option><option>Diyarbakir</option><option>Edirne</option><option>Elazig</option><option>Erzincan</option><option>Erzurum</option><option>Eskisehir</option><option>Gaziantep</option><option>Giresun</option><option>Gumushane</option><option>Hakkari</option><option>Hatay</option><option>Isparta</option><option>Mersin</option><option>Istanbul</option><option>Izmir</option><option>Kars</option><option>Kastamonu</option><option>Kayseri</option><option>Kirklareli</option><option>Kirsehir</option><option>Kocaeli</option><option>Konya</option><option>Kutahya</option><option>Malatya</option><option>Manisa</option><option>Kahramanmaras</option><option>Mardin</option><option>Mugla</option><option>Mus</option><option>Nevsehir</option><option>Nigde</option><option>Ordu</option><option>Rize</option><option>Sakarya</option><option>Samsun</option><option>Siirt</option><option>Sinop</option><option>Sivas</option><option>Tekirdag</option><option>Tokat</option><option>Trabzon</option><option>Tunceli</option><option>Sanliurfa</option><option>Usak</option><option>Van</option><option>Yozgat</option><option>Zonguldak</option><option>Aksaray</option><option>Bayburt</option><option>Karaman</option><option>Kirikkale</option><option>Batman</option><option>Sirnak</option><option>Bartin</option><option>Ardahan</option><option>Igdir</option><option>Yalova</option><option>Karabuk</option><option>Kilis</option><option>Osmaniye</option><option>Duzce</option>
    </select>
    <button type="submit" class="btn btn-primary mt-3">Filtrele</button>
  </form>
</div>
    </header>
    <main>
        <div class="container">
            <div class="row mt-4">
                <div class="col">
                    <table class="table table-hover table-dark table-striped">
                        <thead>
                            <tr>
                                <th>İçerik</th>
                                <th>İl</th>
                                <th>İlçe</th>
                                <th>Tarih</th>
                                <th>Detay/Konum</th>
                            </tr>
                        </thead>
                        <tbody>
                        <?php while($satir = $sorgu->fetch(PDO::FETCH_ASSOC)) { ?>
                            <tr>
                                <td><?=$satir['icerik']?></td>
                                <td><?=$satir['il']?></td>
                                <td><?=$satir['ilce']?></td>
                                <td><?=$satir['date']?></td>
                                <td>
                                    <div class="btn-group">
                                        <a href="detay.php?id=<?=$satir['id']?>" class="btn btn-success">Detay</a>
                                        <a href="https://www.google.com/maps/dir//+<?=$satir['adres']?>" class="btn btn-secondary">Konum</a>
                                    </div>
                                </td>
                            </tr>
                        <?php } ?>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>

    </main>
    <footer><div style="font-size: 80%;" class="mt-5 p-4 bg-dark text-white text-center">
  <p>Bu site yardım amaçlı kurulmuş olup, herhangi bir maddi kazanç gözetmemektedir.
Sitede paylaşılan bilgiler depremzedelerin ve yardımsever vatandaşların birbirleri ile kolay şekilde iletişim kurabilmelerini sağlamak adına depolanacaktır.</p>
</div></footer>
</body>
</html>

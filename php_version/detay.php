
<?php

include "baglan.php";

$sql ="SELECT * FROM yardimlar4 WHERE id = ?";
$sorgu = $baglan->prepare($sql);
$sorgu->execute([
    $_GET['id']
]);
$satir = $sorgu->fetch(PDO::FETCH_ASSOC);
?>
<!DOCTYPE html>
<html lang="tr">
<body style="background-color:#181a1b;">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Yardım Bulunan Bölgeler</title>
    <!-- CSS only -->
<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.1/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-+0n0xVW2eSR5OomGNYDnhzAbDsOXxcvSN1TPprVMTNDbiYZCxYbOOl7+AMvyTG2x" crossorigin="anonymous">
</head>
<body>

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
      &nbsp;
    </header>
    <main>
    <div class="container">
        <div class="row">
            <div class="col">
            <div class="card">
                <div class="card-body">
                  <b>Tarih</b>
                  <p class="card-text"><?=$satir["date"]?></p>
                    <b> İl - İlçe</b>
                    <p class="card-text">
                    <?=$satir["il"]?> <?=$satir["ilce"]?>
                  </p>
                  <b>Açık Adres</b>
                  <p class="card-text"><?=$satir["adres"]?></p>
                    <b>Telefon Numarası</b>
                    <p class="card-text"><?=$satir["tel"]?></p>
                    <b>Yardım İçeriği</b>
                    <p class="card-text"><?=$satir["icerik"]?></p>
                    <b>Açıklama</b>
                    <p class="card-text"><?=$satir["detay"]?></p>
                </div>
                </div>
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

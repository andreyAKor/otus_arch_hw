<?php
require_once __DIR__ . '/vendor/autoload.php';

use classes\handlers\Csv;
use classes\handlers\Json;
use classes\handlers\Txt;
use classes\handlers\Xml;
use classes\reader\ReadFile;
use classes\writer\WriteFile;

if (count($argv) != 3) {
    echo 'Неверный формат запуска. Пример: php ' . $argv[0] . ' in_file.txt out_file.txt' . PHP_EOL;
    exit(1);
}

unset($argv[0]);

try {
    run(...$argv);
} catch (Exception $e) {
    echo 'Исключение: ' . $e->getMessage() . PHP_EOL;
    exit(1);
}

function run(string $inFile, string $outFile)
{
    $xml = new Xml(new ReadFile(), new WriteFile(), $outFile);
    $json = new Json(new ReadFile(), new WriteFile(), $outFile);
    $csv = new Csv(new ReadFile(), new WriteFile(), $outFile);
    $txt = new Txt(new ReadFile(), new WriteFile(), $outFile);

    $xml->setNext($json);
    $json->setNext($csv);
    $csv->setNext($txt);

    $reader = new ReadFile();
    $files = $reader->Read($inFile);

    foreach ($files as $file) {
        $xml->handle(trim($file));
    }
}

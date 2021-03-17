<?php
require_once __DIR__ . '/vendor/autoload.php';

use classes\matrix\Matrix;
use classes\matrix\MatrixAddition;
use classes\reader\ReadFileMatrixes;
use classes\writer\WriteFileMatrixAddition;

if (count($argv) != 3) {
    echo 'Неверный формат запуска. Пример: php ' . $argv[0] . ' F0.txt F1.txt' . PHP_EOL;
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
    $operation = new MatrixAddition(new ReadFileMatrixes($inFile), new WriteFileMatrixAddition($outFile));
    (new Matrix($operation))->Calc();
}


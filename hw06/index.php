<?php
require_once __DIR__ . '/vendor/autoload.php';

use classes\matrix\AbstractTemplate;
use classes\matrix\MatrixGenarator;
use classes\matrix\MatrixDeterminant;
use classes\matrix\TransposeMatrix;
use classes\reader\ReaderInterface;
use classes\reader\ReadFileMatrixes;
use classes\writer\WriteFileMatrixAddition;
use classes\writer\WriterInterface;

if (count($argv) != 4) {
    echo 'Неверный формат запуска. Пример: php ' . $argv[0] . ' in_file.txt out_file.txt transpose' . PHP_EOL;
    exit(1);
}

unset($argv[0]);

try {
    run(...$argv);
} catch (Exception $e) {
    echo 'Исключение: ' . $e->getMessage() . PHP_EOL;
    exit(1);
}

function run(string $inFile, string $outFile, string $operation)
{
    getObjectByOperation(new ReadFileMatrixes($inFile), new WriteFileMatrixAddition($outFile), $operation)->Run();
}

function getObjectByOperation(ReaderInterface $reader, WriterInterface $writer, string $operation): AbstractTemplate
{
    switch ($operation) {
        case 'addition':
            return new MatrixGenarator($reader, $writer);
        case 'transpose':
            return new TransposeMatrix($reader, $writer);
        case 'determinant':
            return new MatrixDeterminant($reader, $writer);
    }

    echo 'Указан не существующий алгоритм' . PHP_EOL;
    exit(1);
}
"""Synthetic, no-download adversarial gates for the cross-media diagnostic.

This never executes original media or imports the uploaded AS3JS package.
"""
from __future__ import annotations

import importlib.util
from copy import deepcopy
import json
from pathlib import Path
import tempfile
import unittest
import warnings
from zipfile import ZIP_STORED, ZipFile, ZipInfo

ROOT = Path(__file__).resolve().parents[1]
spec = importlib.util.spec_from_file_location('media_blocker_census',
                                               ROOT / 'tools/media-blocker-census.py')
assert spec and spec.loader
census = importlib.util.module_from_spec(spec)
spec.loader.exec_module(census)


class MediaCensusBounds(unittest.TestCase):
    def setUp(self) -> None:
        self.dir = tempfile.TemporaryDirectory(prefix='media-census-test-')
        self.addCleanup(self.dir.cleanup)
        self.path = Path(self.dir.name) / 'source.zip'

    def archive(self, members: list[tuple[str, bytes | None]], *, symlink=False) -> None:
        with warnings.catch_warnings():
            warnings.simplefilter('ignore', UserWarning)
            with ZipFile(self.path, 'w', compression=ZIP_STORED) as z:
                for name, contents in members:
                    info = ZipInfo(name)
                    if symlink:
                        info.create_system = 3
                        info.external_attr = 0o120777 << 16
                    z.writestr(info, contents if contents is not None else b'')

    def test_bounded_valid_archive(self) -> None:
        self.archive([('a.txt', b'valid'), ('nested/a.class', b'bytecode')])
        with census.safe_zip(self.path) as z:
            self.assertEqual(z.read('a.txt'), b'valid')

    def test_bad_members_refuse(self) -> None:
        for name in ('../escape', '/absolute', './dot', 'a/../escape',
                     'a//b', 'C:/windows', 'a\\b'):
            with self.subTest(name=name):
                self.archive([(name, b'valid')])
                with self.assertRaisesRegex(ValueError, 'unsafe ZIP member'):
                    census.safe_zip(self.path)

    def test_duplicate_and_symlink_refuse(self) -> None:
        self.archive([('same', b'a'), ('same', b'b')])
        with self.assertRaisesRegex(ValueError, 'unsafe ZIP member'):
            census.safe_zip(self.path)
        self.archive([('link', b'../outside')], symlink=True)
        with self.assertRaisesRegex(ValueError, 'unsafe ZIP member'):
            census.safe_zip(self.path)

    def test_crc_and_size_budgets_refuse(self) -> None:
        self.archive([('a', b'valid-payload')])
        self.path.write_bytes(self.path.read_bytes().replace(b'valid-payload',
                                                           b'corrupt-data!'))
        with self.assertRaisesRegex(ValueError, 'CRC differs'):
            census.safe_zip(self.path)
        self.archive([('a', b'A' * 8)])
        with self.assertRaisesRegex(ValueError, 'total extraction budget'):
            census.safe_zip(self.path, total=7)
        with self.assertRaisesRegex(ValueError, 'total extraction budget'):
            census.safe_zip(self.path, max_entries=0)

    def test_cache_inside_checkout_refuses(self) -> None:
        with self.assertRaisesRegex(ValueError, 'outside'):
            census.inputs(ROOT / 'reports', False)

    def test_pins_disallow_modified_file(self) -> None:
        fake = Path(self.dir.name) / 'Virus.swf'
        fake.write_bytes(b'not the original')
        with self.assertRaisesRegex(ValueError, 'hash/size differs'):
            census.verified(fake, 'Virus.swf')

    def test_separate_original_constructor_checkpoint_is_not_gameplay(self) -> None:
        data = json.loads((ROOT / 'reports/midlet-acotel-constructor.json').read_text())
        linked = census.verify_constructor_checkpoint(data)
        self.assertTrue(linked['unchangedOriginalConstructorExecutedInMagnolia'])
        self.assertTrue(linked['reusablePinnedConstructorAlsoMatched'])
        self.assertTrue(linked['unselectedOriginalCommandVerifiedOnlyByFieldOracle'])
        self.assertFalse(linked['gameStartAppPaintOrInputExecutedInMagnolia'])
        self.assertFalse(linked['thisCensusReexecutedTheConstructor'])
        for path, value in [
            (('magnolia', 'originalGameStartAppExecuted'), True),
            (('magnolia', 'unchangedOriginalConstructorExecuted'), False),
            (('magnolia', 'offscreenArgbSha256'), '0' * 64),
            (('genericPinnedConstructor', 'chargedSteps'), 18928),
            (('genericPinnedConstructor', 'changedJarRefusedBeforeConstruction'), False),
            (('genericPinnedConstructor', 'originalStartAppWorkerPaintOrDeviceInputExecuted'), True),
            (('genericPinnedConstructor', 'requestedDisplayCommandCount'), 1),
            (('reference', 'workerStarted'), True),
            (('comparison', 'independentOriginalGameStartupCompared'), True),
            (('refusals', 'sameSizeZipValidChangedOriginalBeforeConstruction'), False),
        ]:
            wrong = deepcopy(data)
            wrong[path[0]][path[1]] = value
            with self.subTest(path=path), self.assertRaises(ValueError):
                census.verify_constructor_checkpoint(wrong)

    def test_original_direct_key_checkpoint_is_not_device_input(self) -> None:
        data = json.loads((ROOT / 'reports/midlet-acotel-direct-keys.json').read_text())
        constructor = census.verify_constructor_checkpoint(json.loads(
            (ROOT / 'reports/midlet-acotel-constructor.json').read_text()))
        linked = census.verify_direct_key_checkpoint(data, constructor)
        self.assertEqual(linked['callbacks'],
                         ['keyPressed', 'keyReleased', 'keyRepeated'])
        self.assertFalse(linked['MIDPDeviceKeypadInputDelivered'])
        self.assertFalse(linked['originalStartAppPaintOrGameLoopReached'])
        self.assertFalse(linked['thisCensusReexecutedTheDirectMethods'])
        for path, value in [
            (('magnolia', 'originalStartAppExecuted'), True),
            (('magnolia', 'deviceKeypadEventDelivered'), True),
            (('magnolia', 'directKeyTrace'), []),
            (('magnolia', 'offscreenArgbSha256'), '0' * 64),
            (('magnolia', 'chargedStepsIncludingConstructor'),
             constructor['steps']),
            (('reference', 'workerStarted'), True),
            (('comparison', 'independentGameplayOrDeviceEventCompared'), True),
            (('refusals', 'wrongOrderCodeMethodAndExtraCallBeforeEffects'), False),
            (('refusals', 'missingDirectKeyRetentionCapabilityBeforeConstruction'), False),
        ]:
            wrong = deepcopy(data)
            wrong[path[0]][path[1]] = value
            with self.subTest(path=path), self.assertRaises(ValueError):
                census.verify_direct_key_checkpoint(wrong, constructor)

    def test_original_startApp_return_is_not_worker_frame_or_gameplay(self) -> None:
        data = json.loads((ROOT / 'reports/midlet-acotel-start-return.json').read_text())
        constructor = census.verify_constructor_checkpoint(json.loads(
            (ROOT / 'reports/midlet-acotel-constructor.json').read_text()))
        linked = census.verify_start_return_checkpoint(data, constructor)
        self.assertTrue(linked['unchangedOriginalStartAppReturnedInMagnolia'])
        self.assertFalse(linked['nativeWorkerRunPaintVisibleDeviceOrGameplayReached'])
        self.assertTrue(linked['referenceWorkerDrewNonblankOffscreen'])
        self.assertFalse(linked['sameWorkerOrPixelStatesCompared'])
        self.assertFalse(linked['thisCensusReexecutedStartApp'])
        for path, value in [
            (('magnolia', 'originalConstructorAndStartAppExecuted'), False),
            (('magnolia', 'chargedStepsIncludingConstructor'), constructor['steps']),
            (('magnolia', 'originalRunnableRunExecuted'), True),
            (('magnolia', 'originalPaintExecuted'), True),
            (('magnolia', 'deviceDisplaySelected'), True),
            (('magnolia', 'deviceKeypadEventDelivered'), True),
            (('magnolia', 'gameplayOrVisibleFrame'), True),
            (('magnolia', 'offscreenArgbSha256'), '0' * 64),
            (('reference', 'workerThreadRunningInReference'), False),
            (('reference', 'deviceDisplayEventuallySelected'), False),
            (('reference', 'offscreenDistinctColorsAfterWorker'), 1),
            (('comparison', 'sameActualDevicePixels'), True),
            (('comparison', 'sameWorkerExecutionOrVisibleFrame'), True),
            (('refusals', 'unlinkedRunnableAndPaintAndNoContinuation'), False),
        ]:
            wrong = deepcopy(data)
            wrong[path[0]][path[1]] = value
            with self.subTest(path=path), self.assertRaises(ValueError):
                census.verify_start_return_checkpoint(wrong, constructor)

    def test_owned_constructor_is_not_target_startApp_or_input(self) -> None:
        data = json.loads((ROOT / 'reports/midlet-constructor-generic.json').read_text())
        linked = census.verify_owned_constructor_checkpoint(data)
        self.assertTrue(linked['sameOwnedOriginalConstructorReached'])
        self.assertFalse(linked['targetStartAppDirectKeysWorkerOrPaintExecuted'])
        self.assertTrue(linked['independentReferenceCanvasStillUnselected'])
        self.assertFalse(linked['sameDeviceFrameOrGameplayCompared'])
        self.assertFalse(linked['thisCensusReexecutedTheOwnedFixture'])
        for section, key, value in [
            ('magnolia', 'chargedSteps', 47),
            ('magnolia', 'targetStartAppOrDeviceInput', True),
            ('magnolia', 'requestedCanvasClass', 'DirectSliceMIDlet$View'),
            ('reference', 'targetStartAppCalled', True),
            ('comparison', 'sameDeviceFrameOrGameplayCompared', True),
            ('refusals', 'sixWrongHostOptionsAtZeroGuestSteps', False),
            ('unchangedOwnedClassSha256', 'DirectSliceMIDlet.class', '0' * 64),
        ]:
            wrong = deepcopy(data)
            wrong[section][key] = value
            with self.subTest(section=section, key=key), self.assertRaises(ValueError):
                census.verify_owned_constructor_checkpoint(wrong)

    def test_owned_deferred_start_is_not_original_worker_or_frame(self) -> None:
        data = json.loads((ROOT / 'reports/midlet-deferred-generic.json').read_text())
        linked = census.verify_owned_deferred_checkpoint(data)
        self.assertTrue(linked['sameOwnedOriginalStartAppReached'])
        self.assertFalse(linked['MagnoliaWorkerRanOrPainted'])
        self.assertFalse(linked['sameWorkerOrDevicePixelsCompared'])
        self.assertFalse(linked['thisCensusReexecutedTheOwnedFixture'])
        for section, key, value in [
            ('magnolia', 'chargedSteps', 76),
            ('magnolia', 'workerRan', True),
            ('magnolia', 'nativePaintOrActualDeviceSelected', True),
            ('reference', 'realWorkerRan', False),
            ('comparison', 'sameWorkerExecutionOrFrame', True),
            ('refusals', 'fiveWrongHostOptionsAtZeroGuestSteps', False),
            ('unchangedOwnedClassSha256', 'DeferredSliceMIDlet.class', '0' * 64),
        ]:
            wrong = deepcopy(data)
            wrong[section][key] = value
            with self.subTest(section=section, key=key), self.assertRaises(ValueError):
                census.verify_owned_deferred_checkpoint(wrong)

    def test_owned_direct_key_is_not_keypad_or_original_game(self) -> None:
        data = json.loads((ROOT / 'reports/midlet-direct-generic.json').read_text())
        linked = census.verify_owned_direct_checkpoint(data)
        self.assertTrue(linked['sameOwnedConstructorAndDirectKeyBytecodeExecuted'])
        self.assertFalse(linked['MIDPDeviceKeypadOrOriginalGameStartupExecuted'])
        self.assertFalse(linked['sameDeviceInputOrPixelsCompared'])
        self.assertFalse(linked['thisCensusReexecutedTheOwnedFixture'])
        for section, key, value in [
            ('magnolia', 'chargedFinalSteps', 140),
            ('magnolia', 'displayUnselected', False),
            ('magnolia', 'deviceKeyEventsDelivered', True),
            ('magnolia', 'trace', 'keyPressed(50)=1:true:1'),
            ('reference', 'CanvasUnselected', False),
            ('reference', 'deviceKeyEventsDelivered', True),
            ('comparison', 'realDeviceKeypadOrFrameParityCompared', True),
            ('refusals', 'differentAccuratelyPinnedJarCannotUseOriginalClasses', False),
            ('unchangedOwnedClassSha256', 'DirectSliceMIDlet.class', '0' * 64),
        ]:
            wrong = deepcopy(data)
            wrong[section][key] = value
            with self.subTest(section=section, key=key), self.assertRaises(ValueError):
                census.verify_owned_direct_checkpoint(wrong)

    def test_exact_jdk_math_round_checkpoint_is_not_fourier_simulation(self) -> None:
        raw = (ROOT / 'reports/jvm-math-round.json').read_bytes()
        from hashlib import sha256
        self.assertEqual(sha256(raw).hexdigest(), census.MATH_ROUND_REPORT_SHA)
        data = json.loads(raw)
        linked = census.verify_math_round_checkpoint(data)
        self.assertEqual(linked['seededJdkNativeJsRawBitSamples'], 988)
        self.assertFalse(linked['originalFourierSimulationExecutedInMagnolia'])
        self.assertFalse(linked['thisCensusReexecutedTheMathRoundDifferential'])
        for section, key, value in [
            ('portableTargets', 'bundledJS', False),
            ('refusals', 'wrongDescriptorOrInvocationKind', False),
            ('seededRawBits', 'doubleSamples', 591),
            ('FourierStaticOnly', 'originalSimulationExecutedInMagnolia', True),
            ('FourierStaticOnly', 'afterFirstBlockerEntries', 79),
            ('FourierStaticOnly', 'originalJarSha256', '0' * 64),
        ]:
            changed = deepcopy(data)
            changed[section][key] = value
            with self.subTest(section=section, key=key), self.assertRaises(ValueError):
                census.verify_math_round_checkpoint(changed)
        changed = deepcopy(data)
        changed['allJVMOrActionScriptTODOsCompleted'] = True
        with self.assertRaises(ValueError):
            census.verify_math_round_checkpoint(changed)

    def test_as3js_manifest_does_not_claim_a_runtime(self) -> None:
        raw = (ROOT / 'reports/as3-conformance-manifest.json').read_bytes()
        from hashlib import sha256
        self.assertEqual(sha256(raw).hexdigest(), census.AS3_MANIFEST_SHA)
        data = json.loads(raw)
        linked = census.verify_as3_manifest_checkpoint(data)
        self.assertTrue(linked['sourceOnlyDocumentationTaskCompleted'])
        self.assertFalse(linked['anyMagnoliaAS3ABCOrSWFRuntimeGateCompleted'])
        for section, key, value in [
            ('referenceOracles', 'magnoliaAS3ProgramExecuted', True),
            ('maximumInputAndWork', 'magnoliaAS3AcceptedBytes', 4),
            ('emittedTargets', 'abcWriter', True),
            ('referenceArchive', 'selectedSourceFiles', 30),
        ]:
            wrong = deepcopy(data)
            wrong[section][key] = value
            with self.subTest(section=section, key=key), self.assertRaises(ValueError):
                census.verify_as3_manifest_checkpoint(wrong)

    def test_both_original_as3_cws_containers_are_only_envelopes(self) -> None:
        raw = (ROOT / 'reports/swf-as3-cws-current.json').read_bytes()
        from hashlib import sha256
        self.assertEqual(sha256(raw).hexdigest(), census.AS3_CWS_REPORT_SHA)
        data = json.loads(raw)
        linked = census.verify_as3_cws_checkpoint(data)
        self.assertEqual(set(linked), {'SnakeTutorial.swf', 'Virus.swf'})
        self.assertFalse(linked['Virus.swf']['originalABCParsedOrExecuted'])
        self.assertFalse(linked['SnakeTutorial.swf']['thisCensusReexecutedTheCwsOracle'])
        for section, key, value in [
            ('Virus.swf', 'topTags', 304),
            ('Virus.swf', 'ABCParsedOrExecuted', True),
            ('SnakeTutorial.swf', 'framesRendered', True),
            ('SnakeTutorial.swf', 'originalSha256', '0' * 64),
        ]:
            wrong = deepcopy(data)
            wrong['originals'][section][key] = value
            with self.subTest(section=section, key=key), self.assertRaises(ValueError):
                census.verify_as3_cws_checkpoint(wrong)
        wrong = deepcopy(data)
        wrong['originalGamesRenderedInMagnolia'] = True
        with self.assertRaises(ValueError):
            census.verify_as3_cws_checkpoint(wrong)

    def test_original_abc_pool_is_not_avm2_verification(self) -> None:
        raw = (ROOT / 'reports/swf-as3-abc-pools-current.json').read_bytes()
        from hashlib import sha256
        self.assertEqual(sha256(raw).hexdigest(), census.AS3_ABC_POOL_REPORT_SHA)
        data = json.loads(raw)
        linked = census.verify_as3_abc_pool_checkpoint(data)
        self.assertEqual(set(linked), {'SnakeTutorial.swf', 'Virus.swf'})
        self.assertTrue(linked['SnakeTutorial.swf']['originalConstantPoolCountsAndOffsetsCompared'])
        self.assertFalse(linked['Virus.swf']['methodTraitOrOpcodeVerified'])
        self.assertFalse(linked['SnakeTutorial.swf']['thisCensusReexecutedTheAbcPoolOracle'])
        for name, section, key, value in [
            ('Virus.swf', 'pool', 'methodTraitOrOpcodeVerified', True),
            ('Virus.swf', 'pool', 'asciiStringBytes', 3346),
            ('SnakeTutorial.swf', 'pool', 'constantPoolEndOffset', 1112),
            ('SnakeTutorial.swf', 'original', 'extractedAbcSha256', '0' * 64),
        ]:
            wrong = deepcopy(data)
            if section == 'pool':
                wrong['originals'][name]['pool'][key] = value
            else:
                wrong['originals'][name][key] = value
            with self.subTest(name=name, key=key), self.assertRaises(ValueError):
                census.verify_as3_abc_pool_checkpoint(wrong)
        wrong = deepcopy(data)
        wrong['originalGamesExecuted'] = True
        with self.assertRaises(ValueError):
            census.verify_as3_abc_pool_checkpoint(wrong)

    def test_original_abc_method_headers_are_not_traits_or_opcodes(self) -> None:
        raw = (ROOT / 'reports/swf-as3-abc-signatures-current.json').read_bytes()
        from hashlib import sha256
        self.assertEqual(sha256(raw).hexdigest(), census.AS3_ABC_SIGNATURES_REPORT_SHA)
        data = json.loads(raw)
        linked = census.verify_as3_abc_signatures_checkpoint(data)
        self.assertEqual(linked['Virus.swf']['originalMethodSignatureHeadersCompared'], 94)
        self.assertEqual(linked['SnakeTutorial.swf']['postSignaturesUnverifiedBytes'], 1453)
        self.assertFalse(linked['Virus.swf']['metadataClassScriptTraitOrCodeVerified'])
        self.assertFalse(linked['SnakeTutorial.swf']['thisCensusReexecutedTheSignatureOracle'])
        for section, key, value in [
            ('top', 'originalGamesRenderedOrExecuted', True),
            ('Virus.swf', 'metadataClassScriptTraitOrCodeVerified', True),
            ('Virus.swf', 'methodSignatureCount', 95),
            ('SnakeTutorial.swf', 'postSignaturesUnverifiedBytes', 1452),
        ]:
            changed = deepcopy(data)
            if section == 'top':
                changed[key] = value
            else:
                changed['originals'][section]['signatures'][key] = value
            with self.subTest(section=section, key=key), self.assertRaises(ValueError):
                census.verify_as3_abc_signatures_checkpoint(changed)

    def test_original_abc_structural_tables_do_not_verify_opcodes(self) -> None:
        raw = (ROOT / 'reports/swf-as3-abc-structure-current.json').read_bytes()
        from hashlib import sha256
        self.assertEqual(sha256(raw).hexdigest(), census.AS3_ABC_STRUCTURE_REPORT_SHA)
        data = json.loads(raw)
        linked = census.verify_as3_abc_structure_checkpoint(data)
        self.assertEqual(linked['SnakeTutorial.swf']['originalMethodBodyEnvelopesCompared'], 15)
        self.assertEqual(linked['Virus.swf']['originalTraitRecordsStructurallyCompared'], 199)
        self.assertEqual(linked['Virus.swf']['opaqueUnverifiedOpcodeBytes'], 8677)
        self.assertFalse(linked['SnakeTutorial.swf']['AVM2OpcodeSemanticsVerifiedOrGameExecuted'])
        self.assertFalse(linked['Virus.swf']['thisCensusReexecutedTheStructureOracle'])
        for section, key, value in [
            ('top', 'opcodeBytesDecodedOrVerified', 1),
            ('top', 'originalGamesRenderedOrExecuted', True),
            ('opaque', 'acceptedAsExecutableAVM2OrVerifiedCode', True),
            ('Virus.swf', 'opaqueOpcodeBytes', 8676),
            ('SnakeTutorial.swf', 'traitCount', 21),
            ('Virus.swf', 'methodBodyInstructionsOrScopesVerified', True),
        ]:
            wrong = deepcopy(data)
            if section == 'top':
                wrong[key] = value
            elif section == 'opaque':
                wrong['opaqueOpcodeBoundaryDemonstration'][key] = value
            else:
                wrong['originals'][section]['structure'][key] = value
            with self.subTest(section=section, key=key), self.assertRaises(ValueError):
                census.verify_as3_abc_structure_checkpoint(wrong)

    def test_bounded_jdk_selected_archive_does_not_complete_jmol(self) -> None:
        raw = (ROOT / 'reports/jvm-selective-multiclass.json').read_bytes()
        from hashlib import sha256
        self.assertEqual(sha256(raw).hexdigest(), census.JVM_MULTI_CLASS_REPORT_SHA)
        data = json.loads(raw)
        linked = census.verify_selective_multiclass_checkpoint(data)
        self.assertTrue(linked['originalArcLifecycleOnSelectedFourClassJar'])
        self.assertFalse(linked['JmolOriginalGameExecutedOrRendered'])
        self.assertIn('Object.getClass()', linked['JmolFirstBarrierAfterFiveExplicitClasses'])
        self.assertFalse(linked['implicitTransitiveClosureOrFullClient'])
        self.assertFalse(linked['thisCensusReexecutedTheMultiClassOracle'])
        for section, key, value in [
            ('JmolFirstBarrier', 'originalJmolExecuted', True),
            ('JmolFirstBarrier', 'sourceSha256', '0' * 64),
            ('JmolFirstBarrier', 'firstSelectedBarrier', 'whole Jmol runs'),
            ('routes', 'ArcTest', {'JDK11XvfbEqualsNativeOak': True}),
            ('top', 'allOriginalJmolMoleculesOrDevicePixelsRendered', True),
            ('top', 'refusals', []),
        ]:
            wrong = deepcopy(data)
            if section == 'top':
                wrong[key] = value
            else:
                wrong[section][key] = value
            with self.subTest(section=section, key=key), self.assertRaises(ValueError):
                census.verify_selective_multiclass_checkpoint(wrong)

    def test_flash_static_delta_only_one_bitmap_finally(self) -> None:
        current = json.loads((ROOT / 'reports/jvm-flash-after-branching-finally-static-coverage.json').read_text())
        census.verify_flash_static_delta(current)
        changed = deepcopy(current)
        changed['decodedBodies'] += 1
        with self.assertRaisesRegex(ValueError, 'counts changed'):
            census.verify_flash_static_delta(changed)
        changed = deepcopy(current)
        changed['blockers'][0]['message'] += ' fake extra member'
        with self.assertRaisesRegex(ValueError, 'outside the Bitmap'):
            census.verify_flash_static_delta(changed)
        changed = deepcopy(current)
        changed['blockers'][0]['message'] = (
            'com/macromedia/flash/Bitmap.fake()V at byte 8: '
            'bounded legacy jsr/ret: fake other gate')
        with self.assertRaisesRegex(ValueError, 'outside the Bitmap'):
            census.verify_flash_static_delta(changed)


if __name__ == '__main__':
    unittest.main()

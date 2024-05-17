import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:thetiptop_client/src/app_router.dart';
import 'package:thetiptop_client/src/domain/controllers/auth_controller.dart';
import 'package:thetiptop_client/src/domain/controllers/client_controller.dart';
import 'package:thetiptop_client/src/infrastructure/tools/form/validator.dart';
import 'package:thetiptop_client/src/presentation/themes/default_theme.dart';
import 'package:thetiptop_client/src/presentation/widgets/btn/btn_link_widget.dart';
import 'package:thetiptop_client/src/presentation/widgets/btn/btn_action_widget.dart';
import 'package:thetiptop_client/src/presentation/widgets/form/separator_widget.dart';
import 'package:thetiptop_client/src/presentation/widgets/layouts/layout_client_widget.dart';

class SigninScreen extends ConsumerStatefulWidget {
  const SigninScreen({super.key});

  @override
  SigninScreenState createState() => SigninScreenState();
}

class SigninScreenState extends ConsumerState<SigninScreen> {
  // Clé globale pour le widget Form
  final _formKey = GlobalKey<FormState>();

  // État du formulaire
  Map<String, dynamic> formData = {
    'email': '',
    'password': '',
  };

  @override
  Widget build(BuildContext context) {
    final state = ref.watch(clientControllerProvider);
    double screenWidth = MediaQuery.of(context).size.width;

    return LayoutClientWidget(
      child: Form(
        key: _formKey,
        child: Column(
          children: [
            const SizedBox(
              height: DefaultTheme.bigSpacer,
            ),
            TextFormField(
              decoration: const InputDecoration(labelText: 'Adresse email'),
              validator: (value) => Validator().isEmail(value: value),
              onSaved: (value) {
                formData['email'] = value!;
              },
            ),
            const SizedBox(
              height: DefaultTheme.spacer,
            ),
            TextFormField(
              decoration: const InputDecoration(labelText: 'Mot de passe'),
              validator: (value) => Validator().notEmpty(value: value),
              onSaved: (value) {
                formData['password'] = value!;
              },
            ),
            const SizedBox(
              height: DefaultTheme.smallSpacer,
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.end,
              children: [
                BtnLinkWidget(
                  onPressed: () {
                    context.go(AppRouter.passwordRenewRouteName);
                  },
                  text: "Mot de passe oublié",
                ),
              ],
            ),
            const SizedBox(
              height: DefaultTheme.smallSpacer,
            ),
            Row(
              children: [
                BtnActionWidget(
                  isLoading: state.isLoading,
                  onPressed: state.isLoading
                      ? null
                      : () {
                          if (_formKey.currentState!.validate()) {
                            _formKey.currentState!.save();
                            ref.read(authControllerProvider.notifier).signin(formData);
                          }
                        },
                  style: Theme.of(context).outlinedButtonTheme.style,
                  text: "Connexion",
                ),
              ],
            ),
            const SeparatorWidget(
              text: "ou",
            ),
            Row(
              children: [
                BtnActionWidget(
                  onPressed: () {
                    print("ok");
                  },
                  style: Theme.of(context).elevatedButtonTheme.style?.copyWith(
                        backgroundColor: const MaterialStatePropertyAll(
                          DefaultTheme.blueFC,
                        ),
                      ),
                  text: screenWidth > 950 ? "Facebook Connect" : "Facebook\rConnect",
                ),
                const SizedBox(
                  width: DefaultTheme.smallSpacer,
                ),
                BtnActionWidget(
                  onPressed: () {
                    print("ok");
                  },
                  style: Theme.of(context).elevatedButtonTheme.style?.copyWith(
                        backgroundColor: const MaterialStatePropertyAll(
                          DefaultTheme.blueGC,
                        ),
                      ),
                  text: screenWidth > 950 ? "Google Connect" : "Google\rConnect",
                ),
              ],
            ),
            const SizedBox(
              height: DefaultTheme.spacer,
            ),
            MergeSemantics(
              child: Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  const Text(
                    "Vous n'avez pas de compte ? ",
                  ),
                  BtnLinkWidget(
                    onPressed: () {
                      context.go(AppRouter.signupRouteName);
                    },
                    text: "Inscrivez-vous",
                    fontFamily: DefaultTheme.fontFamilyBold,
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
